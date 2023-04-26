package main

import (
	"<%= packageName %>/helper"
	"time"
	"github.com/carlescere/scheduler"
	"runtime"
	"github.com/google/uuid"
	"os"
	"<%= packageName %>/customlogger"
	"github.com/joho/godotenv"

)
func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
  }

/**
Below is the format required by Eureka to register and application instance
{
    "instance": {
        "hostName": "MY_HOSTNAME",
        "app": "org.github.hellosatish.microservicepattern.awesomeproject",
        "vipAddress": "org.github.hellosatish.microservicepattern.awesomeproject",
        "secureVipAddress": "org.github.hellosatish.microservicepattern.awesomeproject"
        "ipAddr": "10.0.0.10",
        "status": "STARTING",
        "port": {"$": "8080", "@enabled": "true"},
        "securePort": {"$": "8443", "@enabled": "true"},
        "healthCheckUrl": "http://WKS-SOF-L011:8080/healthcheck",
        "statusPageUrl": "http://WKS-SOF-L011:8080/status",
        "homePageUrl": "http://WKS-SOF-L011:8080",
        "dataCenterInfo": {
            "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo", 
            "name": "MyOwn"
        },
    }
}
 */

type AppRegistrationBody struct {
	Instance InstanceDetails `json:"instance"`
}

type InstanceDetails struct {
	InstanceId       string         `json:"instanceId"`
	HostName         string         `json:"hostName"`
	App              string         `json:"app"`
	VipAddress       string         `json:"vipAddress"`
	SecureVipAddress string         `json:"secureVipAddress"`
	IpAddr           string         `json:"ipAddr"`
	Status           string         `json:"status"`
	Port             Port           `json:"port"`
	SecurePort       Port           `json:"securePort"`
	HealthCheckUrl   string         `json:"healthCheckUrl"`
	StatusPageUrl    string         `json:"statusPageUrl"`
	HomePageUrl      string         `json:"homePageUrl"`
	DataCenterInfo   DataCenterInfo `json:"dataCenterInfo"`
}
type Port struct {
	Port    string `json:"$"`
	Enabled string `json:"@enabled"`
}

type DataCenterInfo struct {
	Class string `json:"@class"`
	Name  string `json:"name"`
}

// This struct shall be responsible for manager to manage registration with Eureka
type EurekaRegistrationManager struct {
}

func (erm EurekaRegistrationManager) RegisterWithSerivceRegistry(eurekaConfigs RegistrationVariables) {
	customlogger.Printfun("info","Registering service with status : STARTING")    
	body :=  erm.getBodyForEureka("STARTING")    
	helper.MakePostCall(eurekaConfigs.ServiceRegistryURL()+"msg1", body, nil)
	customlogger.Printfun("info","Waiting for 10 seconds for application to start properly")    	
	time.Sleep(10 * time.Second)
	customlogger.Printfun("info","Updating the status to : UP")    	
	bodyUP :=  erm.getBodyForEureka("UP")
	helper.MakePostCall(eurekaConfigs.ServiceRegistryURL()+"msg1", bodyUP, nil)
}

func (erm EurekaRegistrationManager) SendHeartBeat(eurekaConfigs RegistrationVariables) {
	customlogger.Printfun("info","In SendHeartBeat!")    	
	hostname, err := os.Hostname()
	if err != nil{
		customlogger.Printfun("error","Error while getting hostname which shall be used as APP ID")    	
	}
	job := func() {
		customlogger.Printfun("info","sending heartbeat : "+ time.Now().UTC().String())    	
		helper.MakePutCall(eurekaConfigs.ServiceRegistryURL()+"msg1/"+hostname, nil, nil)
	}
	// Run every 25 seconds but not now.
	scheduler.Every(25).Seconds().Run(job)
	runtime.Goexit()

}
func (erm EurekaRegistrationManager) DeRegisterFromServiceRegistry(configs RegistrationVariables) {
	helper.MakePostCall(configs.ServiceRegistryURL(), nil, nil)
}

func (erm EurekaRegistrationManager) getBodyForEureka(status string) *AppRegistrationBody {
	
	instanceId := uuid.New().String() 
	httpport := goDotEnvVariable("SERVICE_PORT")
	hostname, err := os.Hostname()
	if err != nil{
		customlogger.Printfun("error","Enable to find hostname form OS, sending appname as host name")    	
	}

	ipAddress, err := helper.ExternalIP()
	if err != nil{
		customlogger.Printfun("error","Enable to find IP address form OS")    	
	}

	port := Port{httpport,"true"}
	securePort := Port{"8443","false"}
	dataCenterInfo := DataCenterInfo{"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo","MyOwn"}

	homePageUrl := "http://"+hostname+":"+httpport
	statusPageUrl := "http://"+hostname+":"+httpport+"/status"
	healthCheckUrl := "http://"+hostname+":"+httpport+"/healthcheck"

	instance := InstanceDetails{instanceId,hostname, "msg1", "msg1", "msg1",
		ipAddress,status , port,securePort, healthCheckUrl, statusPageUrl, homePageUrl, dataCenterInfo}

	body := &AppRegistrationBody{instance}
	return body
}