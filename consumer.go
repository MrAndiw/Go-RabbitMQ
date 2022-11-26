package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Aplication")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("err connect :")
		log.Println(err)
	}
	defer conn.Close()

	fmt.Println("Successfuly Connect to rabbitMQ instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("err channel :")
		log.Println(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieve : %s \n", d.Body)
		}
	}()

	fmt.Println("Successfuly Connect to rabbitMQ instance")
	fmt.Println("[*] waiting message..")
	<-forever
}
