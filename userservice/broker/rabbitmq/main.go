package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/behnambm/todo/todocommon"
	"github.com/behnambm/todo/userservice/types"
	"github.com/streadway/amqp"
	"log"
)

type UserService interface {
	CreateUser(types.User) (types.User, error)
}

type Broker struct {
	userSvc   UserService
	queueName string
	conn      *amqp.Connection
}

func New(amqpUri, queueName string, userSvc UserService) *Broker {
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpUri)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("[Broker] Successfully connected to RabbitMQ")

	return &Broker{
		queueName: queueName,
		userSvc:   userSvc,
		conn:      connectRabbitMQ,
	}
}

func (b Broker) Listen(ctx context.Context) {
	channelRabbitMQ, channelErr := b.conn.Channel()
	if channelErr != nil {
		log.Fatalln(channelErr)
	}
	defer channelRabbitMQ.Close()

	messages, consumeErr := channelRabbitMQ.Consume(
		b.queueName, // queue name
		"",          // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         // arguments
	)
	if consumeErr != nil {
		log.Panicln("[Broker] HandleMessage -  Consume error: ", consumeErr)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("[Broker] HandleMessage - shutting the listener down...")

			return
		case message, ok := <-messages:
			if !ok {
				log.Println("[Broker] HandleMessage - message channel is closed.")

				return
			}
			go b.HandleMessage(message)
		}
	}
}

func (b Broker) HandleMessage(message amqp.Delivery) {
	switch message.Type {
	case todocommon.MessageTypeUserRegister:
		b.HandleRegisterMessage(message)
	default:
		log.Println("[Broker] HandleMessage - unable to process message with type: ", message.Type)

	}
}

func (b Broker) HandleRegisterMessage(message amqp.Delivery) {
	userMsg := todocommon.UserMessage{}
	err := json.Unmarshal(message.Body, &userMsg)
	if err != nil {
		log.Println("[Broker] HandleRegisterMessage - unable to unmarshal the message: ", err)
		message.Reject(false)

		return
	}

	// create User type out of broker message
	userToBeCreated := types.User{
		Name:     userMsg.Name,
		Email:    userMsg.Email,
		Password: userMsg.Password,
	}

	newUser, userCreateErr := b.userSvc.CreateUser(userToBeCreated)
	if userCreateErr != nil {
		log.Println("[Broker] HandleRegisterMessage - unable create user: ", userCreateErr)
		message.Reject(false)

		return
	}

	fmt.Printf("[Broker] HandleRegisterMessage - user %d created \n", newUser.ID)

	ackErr := message.Ack(false)
	if ackErr != nil {
		log.Println("[Broker] HandleRegisterMessage -  unable to ack: ", ackErr)

		return
	}
}
