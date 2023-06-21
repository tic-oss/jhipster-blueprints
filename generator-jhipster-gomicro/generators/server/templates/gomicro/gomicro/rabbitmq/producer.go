package rabbit 

import(
	"github.com/carlescere/scheduler"
	"github.com/streadway/amqp"
	"github.com/micro/micro/v3/service/logger"
	"os"
)

func Produce(){
	job :=func(){
		rabbitmqurl :=os.Getenv("GO_MICRO_MESSAGE_BROKER")
		conn,err := amqp.Dial(rabbitmqurl)
		if err != nil {
			logger.Errorf("Failed Initializing Broker Connection")
			panic(err)
		}
		channel, err := conn.Channel()
		if err != nil {
			logger.Errorf(err.Error())
		}
		defer channel.Close()
	channel.QueueDeclare(
			"TestQueue", 
			false, 
			false,  
			false,  
			false,   
			nil,    
		)
	err = channel.Publish(
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
}