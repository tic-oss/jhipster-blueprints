package eureka

import (
	"<%= packageName %>/eurekaregistry/helper"
	"time"
	"github.com/carlescere/scheduler"
	"runtime"
	"os"
	"github.com/micro/micro/v3/service/logger"
)
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
	logger.Infof("Registering service with status : STARTING")    
	body :=  erm.getBodyForEureka("STARTING", eurekaConfigs)    
	helper.MakePostCall(eurekaConfigs.ServiceRegistryURL+"<%= baseName %>", body, nil)
	logger.Infof("Waiting for 10 seconds for application to start properly")    	
	time.Sleep(10 * time.Second)
	logger.Infof("Updating the status to : UP")    	
	bodyUP :=  erm.getBodyForEureka("UP", eurekaConfigs)
	helper.MakePostCall(eurekaConfigs.ServiceRegistryURL+"<%= baseName %>", bodyUP, nil)
}

func (erm EurekaRegistrationManager) SendHeartBeat(eurekaConfigs RegistrationVariables) {
	logger.Infof("info","In SendHeartBeat!")    	
	job := func() {
		logger.Infof("sending heartbeat : "+ time.Now().UTC().String())    	
		helper.MakePutCall(eurekaConfigs.ServiceRegistryURL+ "<%= baseName %>/"+eurekaConfigs.InstanceId, nil, nil)
	}
	// Run every 25 seconds but not now.
	scheduler.Every(25).Seconds().Run(job)
	runtime.Goexit()

}
func (erm EurekaRegistrationManager) DeRegisterFromServiceRegistry(configs RegistrationVariables) {
	helper.MakePostCall(configs.ServiceRegistryURL, nil, nil)
}

func (erm EurekaRegistrationManager) getBodyForEureka(status string, configs RegistrationVariables) *AppRegistrationBody {
	httpport := os.Getenv("SERVICE_PORT")
	// hostname, err := os.Hostname()
	// if err != nil{
	// 	logger.Infof("error","Enable to find hostname form OS, sending appname as host name")    	
	// }
	hostname := "<%= baseName %>"

	ipAddress, err := helper.ExternalIP()
	if err != nil{
		logger.Errorf("Enable to find IP address form OS")    	
	}

	port := Port{httpport,"true"}
	securePort := Port{"8443","false"}
	dataCenterInfo := DataCenterInfo{"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo","MyOwn"}

	homePageUrl := "http://"+hostname+":"+httpport
	statusPageUrl := "http://"+hostname+":"+httpport+"/status"
	healthCheckUrl := "http://"+hostname+":"+httpport+"/healthcheck"

	instance := InstanceDetails{configs.InstanceId, hostname, hostname, hostname, hostname,
		ipAddress,status , port,securePort, healthCheckUrl, statusPageUrl, homePageUrl, dataCenterInfo}

	body := &AppRegistrationBody{instance}
	return body
}