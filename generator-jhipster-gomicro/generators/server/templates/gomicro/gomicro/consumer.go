package main

import (
	"github.com/micro/micro/v3/service/logger"
)

func consume() {
    
	channel, err := conn.Channel()
	if err != nil {
		logger.Errorf(err.Error())
	}
	defer channel.Close()

	msgs, err := channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
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