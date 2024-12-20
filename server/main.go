package main

import (
	"context"
	"crypto/tls"
	"fmt"
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
	cfg.Load()
	log.Println(cfg)
	s := service.New(cfg)

	router := gin.Default()

	// 80 -> 443
	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
		})); err != nil {
			log.Fatalf("Failed to start HTTP server: %s", err)
		}
	}()

	// front end webpage
	router.StaticFile("/", cfg.Webpage.BasePath+"/index.html")
	router.StaticFile("/css/index.css", cfg.Webpage.BasePath+"/css/index.css")
	router.StaticFile("/user/", cfg.Webpage.BasePath+"/user.html")
	router.StaticFile("/css/user.css", cfg.Webpage.BasePath+"/css/user.css")
	router.StaticFile("/login/", cfg.Webpage.BasePath+"/login.html")
	router.StaticFile("/css/login.css", cfg.Webpage.BasePath+"/css/login.css")
	router.StaticFile("/password/", cfg.Webpage.BasePath+"/password.html")
	router.StaticFile("/css/password.css", cfg.Webpage.BasePath+"/css/password.css")
	router.StaticFile("/settings/", cfg.Webpage.BasePath+"/settings.html")
	router.StaticFile("/css/settings.css", cfg.Webpage.BasePath+"/css/settings.css")
	router.StaticFile("/summary/", cfg.Webpage.BasePath+"/summary.html")
	router.StaticFile("/css/summary.css", cfg.Webpage.BasePath+"/css/summary.css")
	router.StaticFile("/emphasis/", cfg.Webpage.BasePath+"/emphasis.html")
	router.StaticFile("/css/emphasis.css", cfg.Webpage.BasePath+"/css/emphasis.css")
	// router.GET("/", s.RootHandler)

	api := router.Group("/api")
	{
		// Email service
		api.POST("/summarize", s.Summarize)
		api.POST("/emphasize", s.Emphasize)

		// DB service
		api.POST("/user", s.CreateUser)
		api.PUT("/user/password", s.ChangePwd)
		api.POST("/auth/login", s.Login)
		api.POST("/auth/logout", s.Logout)
		api.GET("/user/settings", s.GetSettings)
		api.PUT("/user/settings", s.UpdateSettings)
		api.GET("/user/summary", s.GetSummary)
		api.GET("/user/emphasis", s.GetEmphasis)
	}

	// https
	cert, err := tls.LoadX509KeyPair(fmt.Sprintf("/etc/letsencrypt/live/%v/fullchain.pem", cfg.Hostname), fmt.Sprintf("/etc/letsencrypt/live/%v/privkey.pem", cfg.Hostname))
	if err != nil {
		log.Fatalf("Failed to load certificate and key: %s", err)
	}

	srv := &http.Server{
		Addr:    ":443",
		Handler: router.Handler(),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
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
