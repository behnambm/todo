package main

import (
	"context"
	"github.com/behnambm/todo/gatewayservice/adapter/grpcadapter"
	"github.com/behnambm/todo/gatewayservice/adapter/rabbitmqadapter"
	"github.com/behnambm/todo/gatewayservice/server/httpserver"
	"github.com/behnambm/todo/gatewayservice/service/authservice"
	"github.com/behnambm/todo/gatewayservice/service/todoservice"
	"github.com/behnambm/todo/gatewayservice/service/userservice"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	HttpListenPort     = os.Getenv("HTTP_LISTEN_PORT")
	AMQPServerUri      = os.Getenv("AMQP_SERVER_URI")
	BrokerUserQueue    = os.Getenv("BROKER_USER_QUEUE")
	BrokerTodoQueue    = os.Getenv("BROKER_TODO_QUEUE")
	UserGRPCServiceURL = os.Getenv("USER_SERVICE_GRPC_URL")
	AuthGRPCServiceURL = os.Getenv("AUTH_SERVICE_GRPC_URL")
	TodoGRPCServiceURL = os.Getenv("TODO_SERVICE_GRPC_URL")
)

func main() {
	checkEnvs()

	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// this goroutine will wait for interrupt signal and stop the service accordingly
	go func() {
		select {
		case s := <-sigCh:
			log.Printf("got signal %v, attempting graceful shutdown", s)
			cancel()
		}
	}()

	gRPCAdapter := grpcadapter.New(UserGRPCServiceURL, AuthGRPCServiceURL, TodoGRPCServiceURL)
	rabbitmqAdapter := rabbitmqadapter.New(AMQPServerUri, BrokerUserQueue, BrokerTodoQueue)

	authService := authservice.New(rabbitmqAdapter, gRPCAdapter)
	userService := userservice.New(rabbitmqAdapter, gRPCAdapter)
	todoService := todoservice.New(rabbitmqAdapter, gRPCAdapter)
	server := httpserver.New(":2020", authService, userService, todoService)

	server.Run(ctx)
}

func checkEnvs() {
	if HttpListenPort == "" {
		log.Fatalf("invalid HTTP_LISTEN_PORT ")
	}
	if AMQPServerUri == "" {
		log.Fatalf("invalid AMQP_SERVER_URI ")
	}
	if BrokerUserQueue == "" {
		log.Fatalf("invalid BROKER_USER_QUEUE ")
	}
	if BrokerTodoQueue == "" {
		log.Fatalf("invalid BROKER_TODO_QUEUE ")
	}
	if UserGRPCServiceURL == "" {
		log.Fatalf("invalid USER_SERVICE_GRPC_URL ")
	}
	if AuthGRPCServiceURL == "" {
		log.Fatalf("invalid AUTH_SERVICE_GRPC_URL ")
	}
	if TodoGRPCServiceURL == "" {
		log.Fatalf("invalid TODO_SERVICE_GRPC_URL ")
	}
}
