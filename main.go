package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {

	// === connect to rabbitMQ ===
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("err connect :")
		log.Println(err)
	}
	defer conn.Close()

	fmt.Println("Successfuly Connect to rabbitMQ instance")

	// === connect to rabbitMQ channel ===
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("err channel :")
		log.Println(err)
	}
	defer ch.Close()

	// === declare queue ===
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("err queue :")
		log.Println(err)
	}

	// === send data to rabbitMQ struct ===
	type Person struct {
		Name string
		Age  int
	}

	Andi := Person{
		Name: "Andi",
		Age:  26,
	}
	AndiJson, _ := json.Marshal(Andi)

	// === publish queue to rabbitMQ ===
	if err = ch.Publish(
		"",     // publish to an exchange
		q.Name, // exchange namerouting to 0 or more queues
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(AndiJson),
		},
	); err != nil {
		fmt.Println("err publish :")
		log.Println(err)
	}

	fmt.Println("Successfuly publish message to queue.")

}
