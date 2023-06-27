package main

import (
	<%_ if (auth){  _%>
	auth "<%= packageName %>/auth"
	<%_ } _%>
	<%_ if (postgresql){  _%>
	"<%= packageName %>/handler"
	<%_ } _%>
	"os"
	"github.com/asim/go-micro/v3"
	<%_ if (rabbitmq){  _%>
	rabbitmq "<%= packageName %>/rabbitmq"
	<%_ } _%>
	"github.com/micro/micro/v3/service/logger"
	<%_ if (eureka){  _%>
	eureka "<%= packageName %>/eurekaregistry"
	"github.com/google/uuid"
	<%_ } _%>
	_ "github.com/jackc/pgx/v4/stdlib"
	"net/http"
	mhttp "github.com/go-micro/plugins/v3/server/http"
   "github.com/joho/godotenv"
   "github.com/gorilla/mux"
   "encoding/json"
)

<%_ if (eureka){  _%>
var configurations eureka.RegistrationVariables
<%_ } _%>

func init(){
	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Error loading .env file")
	}
}

func main() {
	<%_ if (eureka){  _%>
	defer cleanup()
	<%_ } _%>
	<%_ if (auth){  _%>
	auth.SetClient()
	<%_ } _%>
	<%_ if (postgresql){  _%>
	handler.InitializeDb()
	<%_ } _%>
	<%_ if (eureka){  _%>
	service_registry_url :=os.Getenv("GO_MICRO_SERVICE_REGISTRY_URL")
	InstanceId := "<%= baseName %>:"+uuid.New().String()
	configurations = eureka.RegistrationVariables {ServiceRegistryURL:service_registry_url,InstanceId:InstanceId}
	<%_ } _%>
	port :=os.Getenv("GO_MICRO_SERVICE_PORT")
	srv := micro.NewService(
		micro.Server(mhttp.NewServer()),
    )
	opts1 := []micro.Option{
		micro.Name("<%= baseName %>"),
		micro.Version("latest"),
		micro.Address(":"+port),
	}
	srv.Init(opts1...)
	r := mux.NewRouter()
	<%_ if (postgresql){  _%>
	r.Handle("/event",auth.Protect(http.HandlerFunc(handler.CreateHandler))).Methods(http.MethodPost)
	r.Handle("/events",auth.Protect(http.HandlerFunc(handler.ReadHandler))).Methods(http.MethodGet)
	r.Handle("/events/{id}",auth.Protect(http.HandlerFunc(handler.ReadByIdHandler))).Methods(http.MethodGet)
	r.Handle("/update",auth.Protect(http.HandlerFunc(handler.UpdateHandler))).Methods(http.MethodPatch)
	r.Handle("/delete/{id}",auth.Protect(http.HandlerFunc(handler.DeleteHandler))).Methods(http.MethodDelete)
	<%_ } _%>
	r.HandleFunc("/management/health/readiness", func(w http.ResponseWriter, _ *http.Request) {   
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP","components":map[string]interface{} {"readinessState": map[string]interface{}{"status": "UP"}}})}).Methods(http.MethodGet)
	r.HandleFunc("/hello",func(w http.ResponseWriter, _ *http.Request){
		json.NewEncoder(w).Encode("helloworld")}).Methods(http.MethodGet)
	r.HandleFunc("/management/health/liveness", func(w http.ResponseWriter, _ *http.Request) {     
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP","components":map[string]interface{} {"livenessState": map[string]interface{}{"status": "UP"}}})}).Methods(http.MethodGet)
		
	var handlers http.Handler = r
	
	<%_ if (eureka){  _%>
	go eureka.ManageDiscovery(configurations)
	<%_ } _%>

	<%_ if (rabbitmq){  _%>
	rabbitmq.Produce() 
	go rabbitmq.Consume()
    <%_ } _%>
	  
    if err := micro.RegisterHandler(srv.Server(), handlers); err != nil {
		logger.Fatal(err)
	}
	
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}							

func cleanup(){
	eureka.Cleanup(configurations)
}