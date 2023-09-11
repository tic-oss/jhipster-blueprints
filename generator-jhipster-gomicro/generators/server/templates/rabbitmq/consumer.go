package rabbitmq

import (
	"github.com/micro/micro/v3/service/logger"
	"github.com/streadway/amqp"
	app "<%= packageName %>/config"
)

func Consumer<%= rabbitmqServer %>To<%= rabbitmqClient %> () {
    rabbitmqUrl :=app.GetVal("GO_MICRO_MESSAGE_BROKER")
	conn,err := amqp.Dial(rabbitmqUrl)

	exchangeName := "<%= rabbitmqServer %>To<%= rabbitmqClient %>_message_exchange"
	routingKey := "<%= rabbitmqServer %>To<%= rabbitmqClient %>_message_routingKey"
	queueName := "<%= rabbitmqServer %>To<%= rabbitmqClient %>_message_queue"
	
	if err != nil {
		logger.Errorf("Failed Initializing Broker Connection")
	}
	channel, err := conn.Channel()
	if err != nil {
		logger.Errorf(err.Error())
	}
	defer channel.Close()

	msgs, err := channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

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

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			logger.Infof("Recieved Message: %s\n", d.Body)
		}
	}()
	<-forever
}