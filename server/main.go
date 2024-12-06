package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abdtyx/Optimail/server/config"
	"github.com/abdtyx/Optimail/server/service"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialization
	cfg := config.New()
	s := service.New(cfg)

	router := gin.Default()
	router.GET("/", s.RootHandler)

	// Email service
	router.POST("/summarize", s.Summarize)
	router.POST("/emphasize", s.Emphasize)

	// DB service
	router.POST("/user", s.CreateUser)
	router.PUT("/user/:id/password", s.ChangePwd)
	router.POST("/auth/login", s.Login)
	router.POST("/auth/logout", s.Logout)
	router.GET("/user/:id/settings", s.GetSettings)
	router.PUT("/user/:id/settings", s.UpdateSettings)
	router.POST("/user/secret/reset", s.ResetSecret)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

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
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	s.Close()

	log.Println("Server exiting")
}
