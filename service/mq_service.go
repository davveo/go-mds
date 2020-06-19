package service

import "fmt"

type RabbitMQ struct {
}

func (rbq *RabbitMQ) Send(messageBody, messageQueue string) error {
	fmt.Println(messageBody, messageQueue)
	return nil
}
