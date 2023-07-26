package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/behnambm/todo/todocommon"
	"github.com/behnambm/todo/todoservice/types"
	"github.com/streadway/amqp"
	"log"
)

type TodoService interface {
	CreateTodo(types.Todo) (types.Todo, error)
	UpdateTodo(types.Todo) (types.Todo, error)
	DeleteTodo(int64) error
	CreateItem(types.Item) (types.Item, error)
	UpdateItem(types.Item) (types.Item, error)
	DeleteItem(int64) error
}

type Broker struct {
	todoSvc   TodoService
	queueName string
	conn      *amqp.Connection
}

func New(amqpUri, queueName string, todoService TodoService) *Broker {
	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpUri)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully connected to RabbitMQ")

	return &Broker{
		queueName: queueName,
		todoSvc:   todoService,
		conn:      connectRabbitMQ,
	}

}

// Listen starts the RabbitMQ message listener to consume messages from the specified queue
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
		log.Panicln("MQ Consume error: ", consumeErr)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("shutting the listener down...")

			return
		case message, ok := <-messages:
			if !ok {
				log.Println("MQ message channel is closed.")

				return
			}
			go b.HandleMessage(message)
		}
	}
}

// HandleMessage is used to call the appropriate handler based on the message typ
func (b Broker) HandleMessage(message amqp.Delivery) {
	switch message.Type {
	case todocommon.MessageTypeTodoCreate:
		b.HandleTodoCreate(message)
	case todocommon.MessageTypeTodoUpdate:
		b.HandleTodoUpdate(message)
	case todocommon.MessageTypeItemCreate:
		b.HandleItemCreate(message)
	case todocommon.MessageTypeItemUpdate:
		b.HandleItemUpdate(message)
	case todocommon.MessageTypeTodoDelete:
		b.HandleTodoDelete(message)
	case todocommon.MessageTypeItemDelete:
		b.HandleItemDelete(message)
	default:
		log.Println("[Broker] unable to process message with type: ", message.Type)
	}
}

func (b Broker) HandleTodoCreate(message amqp.Delivery) {
	todoMsg := todocommon.TodoMessage{}
	if err := json.Unmarshal(message.Body, &todoMsg); err != nil {
		log.Println("[Broker] HandleTodoCreate - unable to unmarshal the message: ", err)
		// Reject the message since it couldn't be processed correctly & do not queue it again
		message.Reject(false)

		return
	}

	// create type out of broker message
	todoToBeCreated := types.Todo{
		Name:        todoMsg.Name,
		Description: todoMsg.Description,
		UserId:      todoMsg.UserId,
	}

	newTodo, todoUpdateErr := b.todoSvc.CreateTodo(todoToBeCreated)
	if todoUpdateErr != nil {
		log.Println("[Broker] HandleTodoCreate - unable create todo: ", todoUpdateErr)
		message.Reject(false)

		return
	}

	fmt.Printf("[Broker] HandleTodoCreate - todo %d created \n", newTodo.ID)

	ackErr := message.Ack(false)
	if ackErr != nil {
		log.Println("[Broker] HandleTodoCreate - unable to ack: ", ackErr)

		return
	}
}

func (b Broker) HandleTodoUpdate(message amqp.Delivery) {
	todoMsg := todocommon.TodoMessage{}
	if err := json.Unmarshal(message.Body, &todoMsg); err != nil {
		log.Println("[Broker] HandleTodoUpdate - unable to unmarshal the message: ", err)
		message.Reject(false)

		return
	}

	toBeUpdatedTodo := types.Todo{
		Name:        todoMsg.Name,
		Description: todoMsg.Description,
		ID:          todoMsg.ID,
	}

	newTodo, todoUpdateErr := b.todoSvc.UpdateTodo(toBeUpdatedTodo)
	if todoUpdateErr != nil {
		log.Println("[Broker] HandleTodoUpdate - unable update todo: ", todoUpdateErr)
		message.Reject(false)

		return
	}

	fmt.Printf("[Broker] HandleTodoUpdate - todo %d updated \n", newTodo.ID)

	ackErr := message.Ack(false)
	if ackErr != nil {
		log.Println("[Broker] HandleTodoUpdate - unable to ack: ", ackErr)

		return
	}
}

func (b Broker) HandleTodoDelete(message amqp.Delivery) {
	todoMsg := todocommon.TodoMessage{}
	if err := json.Unmarshal(message.Body, &todoMsg); err != nil {
		log.Println("[Broker] HandleTodoDelete - unable to unmarshal the message: ", err)
		message.Reject(false)

		return
	}

	todoUpdateErr := b.todoSvc.DeleteTodo(todoMsg.ID)
	if todoUpdateErr != nil {
		log.Println("[Broker] HandleTodoDelete - unable to delete todo: ", todoUpdateErr)
		message.Reject(false)

		return
	}

	fmt.Printf("[Broker] HandleTodoDelete - todo %d deleted \n", todoMsg.ID)

	ackErr := message.Ack(false)
	if ackErr != nil {
		log.Println("[Broker] HandleTodoDelete - unable to ack: ", ackErr)

		return
	}
}

func (b Broker) HandleItemCreate(message amqp.Delivery) {
	itemMsg := todocommon.ItemMessage{}
	if err := json.Unmarshal(message.Body, &itemMsg); err != nil {
		log.Println("[Broker] HandleItemCreate - unable to unmarshal the message: ", err)

		message.Reject(false)
		return
	}

	toBeCreatedItem := types.Item{
		Title:    itemMsg.Title,
		Priority: itemMsg.Priority,
		TodoId:   itemMsg.TodoId,
		UserId:   itemMsg.UserId,
	}

	newItem, itemCreateErr := b.todoSvc.CreateItem(toBeCreatedItem)
	if itemCreateErr != nil {
		log.Println("[Broker] HandleItemCreate - unable create item: ", itemCreateErr)
		message.Reject(false)

		return
	}

	fmt.Printf("[Broker] HandleItemCreate - item %d created \n", newItem.ID)

	ackErr := message.Ack(false)
	if ackErr != nil {
		log.Println("[Broker] HandleItemCreate - unable to ack: ", ackErr)

		return
	}
}

func (b Broker) HandleItemUpdate(message amqp.Delivery) {
	itemMsg := todocommon.ItemMessage{}
	if err := json.Unmarshal(message.Body, &itemMsg); err != nil {
		log.Println("[Broker] HandleItemUpdate - unable to unmarshal the message: ", err)
		message.Reject(false)

		return
	}

	toBeUpdatedItem := types.Item{
		ID:       itemMsg.ID,
		Title:    itemMsg.Title,
		Priority: itemMsg.Priority,
	}

	newItem, itemUpdateErr := b.todoSvc.UpdateItem(toBeUpdatedItem)
	if itemUpdateErr != nil {
		log.Println("[Broker] HandleItemUpdate - unable update item: ", itemUpdateErr)
		message.Reject(false)

		return
	}

	fmt.Printf("[Broker] HandleItemUpdate - item %d updated \n", newItem.ID)

	ackErr := message.Ack(false)
	if ackErr != nil {
		log.Println("[Broker] HandleItemUpdate - unable to ack: ", ackErr)

		return
	}
}

func (b Broker) HandleItemDelete(message amqp.Delivery) {
	itemMsg := todocommon.ItemMessage{}
	if err := json.Unmarshal(message.Body, &itemMsg); err != nil {
		log.Println("[Broker] HandleItemDelete - unable to unmarshal the message: ", err)
		message.Reject(false)

		return
	}

	itemUpdateErr := b.todoSvc.DeleteItem(itemMsg.ID)
	if itemUpdateErr != nil {
		log.Println("[Broker] HandleItemDelete - unable to delete item: ", itemUpdateErr)
		message.Reject(false)

		return
	}

	fmt.Printf("[Broker] HandleItemDelete - item %d deleted \n", itemMsg.ID)

	ackErr := message.Ack(false)
	if ackErr != nil {
		log.Println("[Broker] HandleItemDelete - unable to ack: ", ackErr)

		return
	}
}
