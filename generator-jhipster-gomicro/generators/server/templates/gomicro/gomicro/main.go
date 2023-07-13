package main

import (
	<%_ if (auth){  _%>
	auth "<%= packageName %>/auth"
	<%_ } _%>
	<%_ if (postgresql||mongodb){  _%>
	"<%= packageName %>/handler"
	<%_ } _%>
	"github.com/asim/go-micro/v3"
	<%_ if (rabbitmqClient?.length||rabbitmqServer?.length){  _%>
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
   "github.com/gorilla/mux"
    jsons "encoding/json"
	app "<%= packageName %>/config"
)

<%_ if (eureka){  _%>
var configurations eureka.RegistrationVariables
<%_ } _%>

func main() {
	<%_ if (eureka){  _%>
	defer cleanup()
	<%_ } _%>
	app.Setconfig()
	<%_ if (auth){  _%>
	auth.SetClient()
	<%_ } _%>
	<%_ if (postgresql){  _%>
	handler.InitializeDb()
	<%_ } _%>
	<%_ if (mongodb){  _%>
	handler.InitializeMongoDb()
	<%_ } _%>
	<%_ if (eureka){  _%>
	service_registry_url :=app.GetVal("GO_MICRO_SERVICE_REGISTRY_URL")
	InstanceId := "<%= baseName %>:"+uuid.New().String()
	configurations = eureka.RegistrationVariables {ServiceRegistryURL:service_registry_url,InstanceId:InstanceId}
	<%_ } _%>
	port :=app.GetVal("GO_MICRO_SERVICE_PORT")
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
	<%_ if (mongodb){  _%>
	r.Handle("/event",auth.Protect(http.HandlerFunc(handler.AddEvent))).Methods(http.MethodPost)
	r.Handle("/events",auth.Protect(http.HandlerFunc(handler.GetEvents))).Methods(http.MethodGet)
	r.Handle("/events/{id}",auth.Protect(http.HandlerFunc(handler.ReadEventById))).Methods(http.MethodGet)
	r.Handle("/update",auth.Protect(http.HandlerFunc(handler.UpdateEvent))).Methods(http.MethodPatch)
	r.Handle("/delete/{id}",auth.Protect(http.HandlerFunc(handler.DeleteEvent))).Methods(http.MethodDelete)
	<%_ } _%>
	r.HandleFunc("/management/health/readiness", func(w http.ResponseWriter, _ *http.Request) {   
	jsons.NewEncoder(w).Encode(map[string]interface{}{"status": "UP","components":map[string]interface{} {"readinessState": map[string]interface{}{"status": "UP"}}})}).Methods(http.MethodGet)
	<%_ if (restClient){  _%>
	r.HandleFunc("/api/services/<%= baseName %>",func(w http.ResponseWriter, _ *http.Request){
	logger.Infof("response sent")
	jsons.NewEncoder(w).Encode(map[string]string{"server": "UP"})
	}).Methods(http.MethodGet)
	<%_ } _%>	
	<%_ if (restServer?.length){  for(var i=0;i<restServer.length;i++){_%>
	r.HandleFunc("/api/services/<%= restServer[i] %>",func(w http.ResponseWriter, r *http.Request) { eureka.Client(w,r,"<%= restServer[i] %>")}).Methods(http.MethodGet)
	<%_ }} _%>	
	r.HandleFunc("/hello",func(w http.ResponseWriter, _ *http.Request){
		jsons.NewEncoder(w).Encode("helloworld")}).Methods(http.MethodGet)
	r.HandleFunc("/management/health/liveness", func(w http.ResponseWriter, _ *http.Request) {     
	jsons.NewEncoder(w).Encode(map[string]interface{}{"status": "UP","components":map[string]interface{} {"livenessState": map[string]interface{}{"status": "UP"}}})}).Methods(http.MethodGet)
		
	var handlers http.Handler = r
	
	<%_ if (eureka){  _%>
	go eureka.ManageDiscovery(configurations)
	<%_ } _%>
	<%_ if (rabbitmqClient?.length){  for(var i=0;i<rabbitmqClient.length;i++){
		 capitalizedServer=baseName.charAt(0).toUpperCase()+baseName.slice(1)
		 capitalizedClient=rabbitmqClient[i].charAt(0).toUpperCase()+rabbitmqClient[i].slice(1)
	_%>
	rabbitmq.Producer<%= capitalizedServer %>To<%= capitalizedClient %>()
    <%_ }}_%>
	<%_ if (rabbitmqServer?.length){ for(var i=0;i<rabbitmqServer.length;i++){ 
		 capitalizedServer=rabbitmqServer[i].charAt(0).toUpperCase()+rabbitmqServer[i].slice(1)
		 capitalizedClient=baseName.charAt(0).toUpperCase()+baseName.slice(1)
	_%>
	go rabbitmq.Consumer<%= capitalizedServer %>To<%= capitalizedClient %>() 
	<%_ }} _%>

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