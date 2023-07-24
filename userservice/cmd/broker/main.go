package main

import (
	"encoding/json"
	"github.com/behnambm/todo/common/types/brokertypes"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	// Create a new RabbitMQ connection.
	amqpServerURL := os.Getenv("AMQP_SERVER_URI")
	queueeName := "user.queue"
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		queueeName, // queue name
		true,       // durable
		false,      // auto delete
		false,      // exclusive
		false,      // no wait
		nil,        // arguments
	)
	if err != nil {
		panic(err)
	}

	msg := brokertypes.UserMessage{
		Name:     "behnam",
		Email:    "behnam@gmail.com",
		Password: "behnam1312",
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		log.Println("error marshaling to json: ", err)
		return
	}

	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
		Type:        brokertypes.MessageTypeUserRegister,
	}

	channelRabbitMQ.Publish(
		"",
		queueeName,
		false,
		false,
		message,
	)
}
