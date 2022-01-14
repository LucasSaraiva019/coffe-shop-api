package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Handlers
	product := handlers.NewProducts(l)

	// Create new Server mux and Register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", product)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		l.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieveed terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
