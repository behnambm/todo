package main

import (
	"context"
	"github.com/behnambm/todo/gatewayservice/adapter/grpcadapter"
	"github.com/behnambm/todo/gatewayservice/adapter/rabbitmqadapter"
	"github.com/behnambm/todo/gatewayservice/server/httpserver"
	"github.com/behnambm/todo/gatewayservice/service/authservice"
	"github.com/behnambm/todo/gatewayservice/service/todoservice"
	"github.com/behnambm/todo/gatewayservice/service/userservice"
	"github.com/behnambm/todo/todocommon"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	HttpListenPort     = todocommon.GetEnvOrPanic("HTTP_LISTEN_PORT")
	AMQPServerUri      = todocommon.GetEnvOrPanic("AMQP_SERVER_URI")
	BrokerUserQueue    = todocommon.GetEnvOrPanic("BROKER_USER_QUEUE")
	BrokerTodoQueue    = todocommon.GetEnvOrPanic("BROKER_TODO_QUEUE")
	UserGRPCServiceURL = todocommon.GetEnvOrPanic("USER_SERVICE_GRPC_URL")
	AuthGRPCServiceURL = todocommon.GetEnvOrPanic("AUTH_SERVICE_GRPC_URL")
	TodoGRPCServiceURL = todocommon.GetEnvOrPanic("TODO_SERVICE_GRPC_URL")
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
	server := httpserver.New(":"+HttpListenPort, authService, userService, todoService)

	server.Run(ctx)
}
