package main

import (
	"context"
	"fmt"
	"github.com/behnambm/todo/common/utils"
	"github.com/behnambm/todo/userservice/broker/rabbitmq"
	"github.com/behnambm/todo/userservice/repo/sqliterepo"
	"github.com/behnambm/todo/userservice/server/grpcserver"
	"github.com/behnambm/todo/userservice/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	GRPCListenPort  = utils.GetEnvOrPanic("GRPC_LISTEN_PORT")
	AMQPServerUri   = utils.GetEnvOrPanic("AMQP_SERVER_URI")
	BrokerUserQueue = utils.GetEnvOrPanic("BROKER_USER_QUEUE")
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
	userService := service.New(sqliteRepo)

	server := grpcserver.New(fmt.Sprintf(":%s", GRPCListenPort), userService)
	go server.Run(ctx)

	broker := rabbitmq.New(
		AMQPServerUri,
		BrokerUserQueue,
		userService,
	)
	broker.Listen(ctx)
}
