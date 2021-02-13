package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/Jimeux/app-mesher/svc-client/config"
	"github.com/Jimeux/app-mesher/svc-client/rest/handlers"
	"github.com/Jimeux/app-mesher/svc-client/rest/router"
	"github.com/Jimeux/app-mesher/svc-client/rpc"
)

func main() {
	conf := config.New()

	conn, err := grpc.Dial(conf.Server.IdentityHost,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	identitySvc := rpc.NewIdentityServiceClient(conn)

	handler := router.Init(&router.Handlers{
		Token: handlers.NewTokenHandler(identitySvc),
	})

	server := &http.Server{
		Addr:              ":" + conf.Server.Port,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	gracefulShutdown := make(chan struct{})
	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		// Asynchronously listen for signals
		signal.Notify(shutdownSignal, syscall.SIGINT)  // ctrl+C
		signal.Notify(shutdownSignal, syscall.SIGTERM) // docker stop
		// Block while waiting for shutdownSignal to receive a signal
		<-shutdownSignal

		// Initiate graceful shut down
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(gracefulShutdown)
	}()

	log.Println("Listening on " + server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("unexpected server error: %v", err)
		close(gracefulShutdown)
	}

	<-gracefulShutdown
}
