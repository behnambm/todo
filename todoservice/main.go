package main

import (
	"context"
	"fmt"
	"github.com/behnambm/todo/todocommon"
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
	GRPCListenPort  = todocommon.GetEnvOrPanic("GRPC_LISTEN_PORT")
	AMQPServerUri   = todocommon.GetEnvOrPanic("AMQP_SERVER_URI")
	BrokerTodoQueue = todocommon.GetEnvOrPanic("BROKER_TODO_QUEUE")
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
