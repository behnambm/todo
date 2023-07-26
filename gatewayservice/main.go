package main

import (
	"context"
	"github.com/behnambm/todo/common/utils"
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
	HttpListenPort     = utils.GetEnvOrPanic("HTTP_LISTEN_PORT")
	AMQPServerUri      = utils.GetEnvOrPanic("AMQP_SERVER_URI")
	BrokerUserQueue    = utils.GetEnvOrPanic("BROKER_USER_QUEUE")
	BrokerTodoQueue    = utils.GetEnvOrPanic("BROKER_TODO_QUEUE")
	UserGRPCServiceURL = utils.GetEnvOrPanic("USER_SERVICE_GRPC_URL")
	AuthGRPCServiceURL = utils.GetEnvOrPanic("AUTH_SERVICE_GRPC_URL")
	TodoGRPCServiceURL = utils.GetEnvOrPanic("TODO_SERVICE_GRPC_URL")
)

func main() {
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
