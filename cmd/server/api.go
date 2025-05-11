package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	cfg "github.com/inidaname/mosque/api_gateway/config"
	"github.com/inidaname/mosque/api_gateway/routes"
)

var app = cfg.CreateApplication()

func Run() error {
	routes := routes.SetupRouter()
	// srv := &http.Server{
	// 	Addr:         app.Config.Addr,
	// 	Handler:      routes,
	// 	WriteTimeout: time.Second * 30,
	// 	ReadTimeout:  time.Second * 10,
	// 	IdleTimeout:  time.Minute,
	// }

	// shutdown := make(chan error)
	// go func() {
	// 	quit := make(chan os.Signal, 1)
	// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 	s := <-quit

	// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 	defer cancel()

	// 	app.Logger.Info("signal caught", "signal", s.String())

	// 	shutdown <- srv.Shutdown(ctx)
	// }()

	// app.Logger.Info("Starting server on", "port", app.Config.Addr, "env", app.Config.Env)

	// err := srv.ListenAndServe()
	// if err != nil {
	// 	return err
	// }

	// app.Logger.Info("server has stopped", "addr", app.Config.Addr, "env", app.Config.Env)

	// Create server
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      routes,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine so it doesn't block
	go func() {
		log.Printf("Starting server on port %s", app.Config.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
	app.Logger.Info("server has stopped", "addr", app.Config.Addr, "env", app.Config.Env)
	return nil
}
