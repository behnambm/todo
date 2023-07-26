package rabbitmqadapter

import (
	"encoding/json"
	"fmt"
	"github.com/behnambm/todo/gatewayservice/types"
	"github.com/behnambm/todo/todocommon"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQAdapter struct {
	conn          *amqp.Connection
	userQueueName string
	todoQueueName string
}

func New(amqpServerURI, userQueueName, todoQueueName string) RabbitMQAdapter {
	connectRabbitMQ, err := amqp.Dial(amqpServerURI)
	if err != nil {
		log.Fatalln(err)
	}

	return RabbitMQAdapter{
		conn:          connectRabbitMQ,
		userQueueName: userQueueName,
		todoQueueName: todoQueueName,
	}
}

func (r RabbitMQAdapter) CreateUser(user types.User) error {
	channelRabbitMQ, err := r.conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channelRabbitMQ.Close()

	msg := todocommon.UserMessage{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	msgBody, jsonErr := json.Marshal(msg)
	if jsonErr != nil {
		log.Println("[Broker Adapter] CreateUser - error marshaling to json: ", jsonErr)

		return fmt.Errorf("[Broker Adapter] CreateUser - error marshaling to json - %w", err)
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
		Type:        todocommon.MessageTypeUserRegister,
	}

	publishErr := channelRabbitMQ.Publish("", r.userQueueName, false, false, message)
	if publishErr != nil {
		return fmt.Errorf("[Broker Adapter] CreateUser - %w", err)
	}

	return nil
}

func (r RabbitMQAdapter) CreateTodo(todo types.Todo) error {
	channelRabbitMQ, err := r.conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channelRabbitMQ.Close()

	msg := todocommon.TodoMessage{
		Name:        todo.Name,
		Description: todo.Description,
		UserId:      todo.UserId,
	}

	msgBody, jsonErr := json.Marshal(msg)
	if jsonErr != nil {
		log.Println("[Broker Adapter] CreateTodo - error marshaling to json: ", jsonErr)

		return fmt.Errorf("[Broker Adapter] CreateTodo - error marshaling to json - %w", err)
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
		Type:        todocommon.MessageTypeTodoCreate,
	}

	publishErr := channelRabbitMQ.Publish("", r.todoQueueName, false, false, message)
	if publishErr != nil {
		return fmt.Errorf("[Broker Adapter] CreateTodo - %w", err)
	}

	return nil

}

func (r RabbitMQAdapter) UpdateTodo(todo types.Todo) error {
	channelRabbitMQ, err := r.conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channelRabbitMQ.Close()

	msg := todocommon.TodoMessage{
		ID:          todo.ID,
		Name:        todo.Name,
		Description: todo.Description,
	}

	msgBody, jsonErr := json.Marshal(msg)
	if jsonErr != nil {
		return fmt.Errorf("[Broker Adapter] UpdateTodo - error marshaling to json - %w", err)
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
		Type:        todocommon.MessageTypeTodoUpdate,
	}

	publishErr := channelRabbitMQ.Publish("", r.todoQueueName, false, false, message)
	if publishErr != nil {
		return fmt.Errorf("[Broker Adapter] UpdateTodo - %w", err)
	}

	return nil

}

func (r RabbitMQAdapter) DeleteTodo(todoId int64) error {
	channelRabbitMQ, err := r.conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channelRabbitMQ.Close()

	msg := todocommon.TodoMessage{ID: todoId}

	msgBody, jsonErr := json.Marshal(msg)
	if jsonErr != nil {
		return fmt.Errorf("[Broker Adapter] DeleteTodo - error marshaling to json - %w", err)
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
		Type:        todocommon.MessageTypeTodoDelete,
	}

	publishErr := channelRabbitMQ.Publish("", r.todoQueueName, false, false, message)
	if publishErr != nil {
		return fmt.Errorf("[Broker Adapter] DeleteTodo - %w", err)
	}

	return nil

}

func (r RabbitMQAdapter) CreateItem(item types.Item) error {
	channelRabbitMQ, err := r.conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channelRabbitMQ.Close()

	msg := todocommon.ItemMessage{
		Title:    item.Title,
		Priority: item.Priority,
		TodoId:   item.TodoId,
		UserId:   item.UserId,
	}

	msgBody, jsonErr := json.Marshal(msg)
	if jsonErr != nil {
		log.Println("[Broker Adapter] CreateItem - error marshaling to json: ", jsonErr)

		return fmt.Errorf("[Broker Adapter] CreateItem - error marshaling to json - %w", err)
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
		Type:        todocommon.MessageTypeItemCreate,
	}

	publishErr := channelRabbitMQ.Publish("", r.todoQueueName, false, false, message)
	if publishErr != nil {
		return fmt.Errorf("[Broker Adapter] CreateItem - %w", err)
	}

	return nil

}

func (r RabbitMQAdapter) UpdateItem(item types.Item) error {
	channelRabbitMQ, err := r.conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channelRabbitMQ.Close()

	msg := todocommon.ItemMessage{
		ID:       item.ID,
		Title:    item.Title,
		Priority: item.Priority,
	}

	msgBody, jsonErr := json.Marshal(msg)
	if jsonErr != nil {
		log.Println("[Broker Adapter] UpdateItem - error marshaling to json: ", jsonErr)

		return fmt.Errorf("[Broker Adapter] UpdateItem - error marshaling to json - %w", err)
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
		Type:        todocommon.MessageTypeItemUpdate,
	}

	publishErr := channelRabbitMQ.Publish("", r.todoQueueName, false, false, message)
	if publishErr != nil {
		return fmt.Errorf("[Broker Adapter] UpdateItem - %w", err)
	}

	return nil
}

func (r RabbitMQAdapter) DeleteItem(itemId int64) error {
	channelRabbitMQ, err := r.conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}
	defer channelRabbitMQ.Close()

	msg := todocommon.ItemMessage{ID: itemId}

	msgBody, jsonErr := json.Marshal(msg)
	if jsonErr != nil {
		log.Println("[Broker Adapter] DeleteItem - error marshaling to json: ", jsonErr)

		return fmt.Errorf("[Broker Adapter] DeleteItem - error marshaling to json - %w", err)
	}
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        msgBody,
		Type:        todocommon.MessageTypeItemDelete,
	}

	publishErr := channelRabbitMQ.Publish("", r.todoQueueName, false, false, message)
	if publishErr != nil {
		return fmt.Errorf("[Broker Adapter] DeleteItem - %w", err)
	}

	return nil
}
