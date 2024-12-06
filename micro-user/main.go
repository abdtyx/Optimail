package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abdtyx/Optimail/micro-user/config"
	"github.com/abdtyx/Optimail/micro-user/grpc"
)

func main() {
	cfg := config.New()
	cfg.Load()

	grpc_srv := grpc.NewServer(cfg)

	// service connections
	go grpc_srv.ListenAndServe()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	sigCh := make(chan os.Signal, 1)

	// Gracefully handle signal
	// Registered signals
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	// Waiting for signal
	<-sigCh

	// Shutdown
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := grpc_srv.Shutdown(ctx); err != nil {
		log.Fatal("Grpc Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
