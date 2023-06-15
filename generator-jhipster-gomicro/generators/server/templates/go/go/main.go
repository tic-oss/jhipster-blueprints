package main

import (
	<%_ if (postgresql){  _%>
	"database/sql"
	<%_ } _%>
	pb "<%= packageName %>/proto"
	<%_ if (auth){  _%>
	auth "<%= packageName %>/auth"
	<%_ } _%>
	"context"
	"log"
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
	"github.com/asim/go-micro/v3/broker"
	"github.com/asim/go-micro/plugins/broker/rabbitmq/v3"
	<%_ } _%>
	"github.com/asim/go-micro/v3/server"
	"github.com/micro/micro/v3/service/logger"
	_ "github.com/jackc/pgx/v4/stdlib"
   "github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf("Error loading .env file")
	}
}

<%_ if (rabbitmq){  _%>
func initBroker() broker.Broker {
    // Configure RabbitMQ broker
	rabbitmqurl :=os.Getenv("MESSAGE_BROKER")
	rabbitmqBroker := rabbitmq.NewBroker(
        broker.Addrs(rabbitmqurl),
    )

    if err := rabbitmqBroker.Init(); err != nil {
        log.Fatal(err)
    }
    if err := rabbitmqBroker.Connect(); err != nil {
        log.Fatal(err)
    }

    return rabbitmqBroker
}
<%_ } _%>

<%_ if (mongodb){ _%> 
func GetClient() *mongo.Client {
    clientOptions := options.Client().ApplyURI("mongodb+srv://harsha:harsha@cluster0.l7oje6h.mongodb.net/?retryWrites=true&w=majority")
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Connect(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    return client
}
<%_ } _%>

func main() {
	<%_ if (auth){  _%>
	auth.SetClient()
	<%_ } _%>
	<%_ if (postgresql){  _%>
	dbAddress := os.Getenv("DB_URL")
	<%_ } _%>
	<%_ if (mongodb){ _%> 
    dbAddress := os.Getenv("DB_URL")	
	<%_ } _%>
	<%_ if (eureka){  _%>
	eurekaurl :=os.Getenv("SERVICE_REGISTRY_URL")
	opts := []registry.Option{
		registry.Addrs(eurekaurl),
	}
	<%_ } _%>
	port :=os.Getenv("SERVICE_PORT")
	<%_ if (rabbitmq){  _%>
	broker :=initBroker()
	<%_ } _%>
	srv := micro.NewService(
		micro.Name("<%= baseName %>"),
		micro.Version("latest"),
		micro.Server(
			server.NewServer(
			server.Name("<%= baseName %>"),
			server.Address(":"+string(port)),
		),
	 ),
	 <%_ if (rabbitmq){  _%>
	 micro.Broker(broker),
	 <%_ } _%>
	 <%_ if (eureka){  _%>
	 micro.Registry(eureka.NewRegistry(
         opts...
	)),
	<%_ } _%>
	<%_ if (auth){  _%>
	micro.WrapHandler(auth.AuthWrapper),
	<%_ } _%>
    )
	srv.Init()
	
	<%_ if (postgresql){  _%>
	sqlDB, err := sql.Open("pgx", dbAddress)
	if err != nil {
		logger.Fatalf("Failed to open connection to DB %s", err)
	}
	h := &handler.Db{}
	h.DBConn(sqlDB)

	pb.RegisterDbHandler(srv.Server(), h)
	<%_ } _%>

	<%_ if (mongodb){ _%> 
     c :=GetClient()
	 err := c.Ping(context.Background(), nil)
     if err != nil {
         log.Fatal("Couldn't connect to the database", err)
     } else {
         log.Println("Connected!")
     }
     h := &handler.Db{}
	 h.DBConn(c)
	 pb.RegisterDbHandler(srv.Server(), h)
	<%_ } _%>

	// Subscribe to the topic
	<%_ if (rabbitmq){  _%>
	if err := srv.Server().Subscribe(
		srv.Server().NewSubscriber("my-topic", func(ctx context.Context, message *pb.MyMessage) error {
			// Process the received message
			logger.Infof("Received message: %v",message)
			return nil
		}),
	); err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	// Create a new publisher
	publisher := micro.NewPublisher("my-topic", srv.Client())

	// Publish a message
	job := func(){
		err := publisher.Publish(context.TODO(), &pb.MyMessage{
		Id:"1",
		Data: "Hello, World!",
	});
	if( err != nil) {
		log.Fatalf("Failed to publish message: %v", err)
	}
    }
    scheduler.Every(25).Seconds().Run(job)
	<%_ } _%>    

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}							