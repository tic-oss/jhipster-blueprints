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
	<%_ if (eureka){  _%>
	"github.com/go-micro/plugins/v3/registry/eureka"
	"github.com/asim/go-micro/v3/registry"
	<%_ } _%>
	<%_ if (rabbitmq){  _%>
	"github.com/carlescere/scheduler"
	"github.com/streadway/amqp"
	<%_ } _%>
	"github.com/micro/micro/v3/service/logger"
	_ "github.com/jackc/pgx/v4/stdlib"
	"net/http"
	mhttp "github.com/go-micro/plugins/v3/server/http"
   "github.com/joho/godotenv"
   "github.com/gorilla/mux"
   "encoding/json"
)

<%_ if (rabbitmq){  _%>
var conn *amqp.Connection
<%_ } _%>


func init(){
	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Error loading .env file")
	}
}

<%_ if (rabbitmq){  _%>
	func Initializerabbitmq() *amqp.Connection{
		rabbitmqurl :=os.Getenv("GO_MICRO_MESSAGE_BROKER")
		conn, err := amqp.Dial(rabbitmqurl)
		if err != nil {
			logger.Errorf("Failed Initializing Broker Connection")
			panic(err)
		}
		return conn
	}
<%_ } _%>

func main() {
	<%_ if (auth){  _%>
	auth.SetClient()
	<%_ } _%>
	<%_ if (postgresql){  _%>
	handler.InitializeDb()
	<%_ } _%>
	<%_ if (eureka){  _%>
	eurekaurl :=os.Getenv("GO_MICRO_SERVICE_REGISTRY_URL")
	opts := []registry.Option{
		registry.Addrs(eurekaurl),
	}
	<%_ } _%>
	port :=os.Getenv("GO_MICRO_SERVICE_PORT")
	<%_ if (rabbitmq){  _%>
	conn = Initializerabbitmq()
	<%_ } _%>
	srv := micro.NewService(
		micro.Server(mhttp.NewServer()),
	 <%_ if (eureka){  _%>
	 micro.Registry(eureka.NewRegistry(
         opts...
	)),
	<%_ } _%>
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
	
	r.HandleFunc("/management/health/liveness", func(w http.ResponseWriter, _ *http.Request) {     
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP","components":map[string]interface{} {"livenessState": map[string]interface{}{"status": "UP"}}})}).Methods(http.MethodGet)
		
	var handlers http.Handler = r
    
	<%_ if (rabbitmq){  _%>
	go consume() 
	channel, err := conn.Channel()
	if err != nil {
		logger.Errorf(err.Error())
	}
	defer channel.Close()

	job :=func(){
	err := channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)
	
	if err != nil {
		logger.Errorf(err.Error())
	}
    logger.Infof("Published Message to Queue")
   }

   scheduler.Every(25).Seconds().Run(job)
    <%_ } _%>
	  
    if err := micro.RegisterHandler(srv.Server(), handlers); err != nil {
		logger.Fatal(err)
	}
	
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}							