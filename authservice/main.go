package main

import (
	"context"
	"fmt"
	"github.com/behnambm/todo/authservice/server/grpcserver"
	"github.com/behnambm/todo/authservice/service"
	"github.com/behnambm/todo/todocommon"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	GRPCListenPort = todocommon.GetEnvOrPanic("GRPC_LISTEN_PORT")
	JWTSecret      = todocommon.GetEnvOrPanic("JWT_SECRET")
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

			// cancel the context
			cancel()
		}
	}()

	tokenService := service.New(JWTSecret)
	server := grpcserver.New(fmt.Sprintf(":%s", GRPCListenPort), tokenService)
	server.Run(ctx)
}
