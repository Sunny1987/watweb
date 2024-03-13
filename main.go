package main

import (
	"context"
	"github.com/common-nighthawk/go-figure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"webapp/server"
)

func main() {

	//app logger
	l := log.New(os.Stdout, "WebApp", log.LstdFlags)

	//set router handler
	routerHandler := server.NewAppLogger(l)

	//set server mux
	mux := http.NewServeMux()

	//Register the handlers to the server mux
	mux.HandleFunc("/", routerHandler.StartApp)
	mux.HandleFunc("/scan", routerHandler.StartScan)

	port := ":8000"

	//Define server
	prodServer := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 50 * time.Minute,
		IdleTimeout:  50 * time.Minute,
		ErrorLog:     l,
	}

	go func() {
		myFigure := figure.NewFigure("WATWEB:", "", true)
		myFigure.Print()
		l.Println("version: 1.0.0")
		l.Println("Author: Sabyasachi Roy")
		l.Println("Starting server...")
		if err := prodServer.ListenAndServe(); err != nil {
			l.Printf("Error starting server: %v", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan

	l.Println("Stopping server as per user interrupt", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := prodServer.Shutdown(tc)
	if err != nil {
		return
	}

}
