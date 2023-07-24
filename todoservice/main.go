package main

import (
	"context"
	"fmt"
	"github.com/behnambm/todo/todoservice/broker/rabbitmq"
	"github.com/behnambm/todo/todoservice/repo/sqliterepo"
	"github.com/behnambm/todo/todoservice/server/grpcserver"
	"github.com/behnambm/todo/todoservice/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	GRPCListenPort  = os.Getenv("GRPC_LISTEN_PORT")
	AMQPServerUri   = os.Getenv("AMQP_SERVER_URI")
	BrokerTodoQueue = os.Getenv("BROKER_TODO_QUEUE")
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

	sqliteRepo := sqliterepo.New("storage.db")
	todoService := service.New(sqliteRepo)

	server := grpcserver.New(fmt.Sprintf(":%s", GRPCListenPort), todoService)
	go server.Run(ctx)

	broker := rabbitmq.New(
		AMQPServerUri,
		BrokerTodoQueue,
		todoService,
	)
	broker.Listen(ctx)
}

func checkEnvs() {
	if GRPCListenPort == "" {
		log.Fatalf("invalid GRPC_LISTEN_PORT ")
	}
	if AMQPServerUri == "" {
		log.Fatalf("invalid AMQP_SERVER_URI ")
	}
	if BrokerTodoQueue == "" {
		log.Fatalf("invalid BROKER_TODO_QUEUE ")
	}
}
