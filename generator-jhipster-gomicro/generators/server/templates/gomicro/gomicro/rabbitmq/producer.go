package rabbitmq

import (
	"github.com/carlescere/scheduler"
	"github.com/streadway/amqp"
	"github.com/micro/micro/v3/service/logger"
	app "<%= packageName %>/config"
	"time"
	"encoding/json"
)

func Producer<%= rabbitmqServer %>To<%= rabbitmqClient %>() {
	var counter int 
	job := func() {
		counter++
		rabbitmqUrl := app.GetVal("GO_MICRO_MESSAGE_BROKER")
		conn, err := amqp.Dial(rabbitmqUrl)
		if err != nil {
			logger.Errorf("Failed Initializing Broker Connection")
			panic(err)
		}
		channel, err := conn.Channel()
		if err != nil {
			logger.Errorf(err.Error())
		}
		defer channel.Close()

		exchangeName := "<%= rabbitmqServer %>To<%= rabbitmqClient %>_message_exchange"
		routingKey := "<%= rabbitmqServer %>To<%= rabbitmqClient %>_message_routingKey"
		queueName := "<%= rabbitmqServer %>To<%= rabbitmqClient %>_message_queue"

		err = channel.ExchangeDeclare(
			exchangeName, 
			"direct",    
			true,       
			false,    
			false,     
			false,      
			nil,          
		)
		if err != nil {
			logger.Errorf(err.Error())
		}

		_, err = channel.QueueDeclare(
			queueName, 
			false,     
			false,    
			false,    
			false,    
			nil,       
		)
		if err != nil {
			logger.Errorf(err.Error())
		}

		err = channel.QueueBind(
			queueName,    
			routingKey,    
			exchangeName,  
			false,         
			nil,           
		)
		if err != nil {
			logger.Errorf(err.Error())
		}

		data:=map[string]interface{}{"id":counter,"message":"Publishing this message from <%= baseName %> with key: "+queueName,"dateTime":((time.Now().UnixNano())/int64(time.Millisecond))}
		payload, err := json.Marshal(data)
		
		if err != nil {
		logger.Errorf("Error:", err)
		return
		}

		err = channel.Publish(
			exchangeName,
			routingKey,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body: payload,
			},
		)

		if err != nil {
			logger.Errorf(err.Error())
		}
		logger.Infof("Message Published Successfully")
	}

	scheduler.Every(25).Seconds().Run(job)
}
