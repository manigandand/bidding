package main

import (
	v1 "bidding/api/v1"
	"bidding/config"
	appmiddleware "bidding/middleware"
	"bidding/pkg/trace"
	"flag"
	"log"
	"time"

	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

var (
	// Port in which the bidder client to run
	Port int
	// Delay after the bidder to respond to auction request
	Delay time.Duration
)

func main() {
	name := flag.String("name", "Batman", "name identifier")
	port := flag.Int("port", 0, "port in which the bidder client to run")
	delay := flag.Uint("delay", 0, "delay after the bidder to respond to auction request")

	flag.Parse()

	if *port == 0 {
		log.Fatal("invalid port or port required")
	}
	if *delay > 500 {
		log.Println("delay more than 500ms 😟")
	}
	Port = *port
	Delay = (time.Duration(*delay) * time.Millisecond)

	trace.Setup(config.Env)
	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// cross & loger middleware
	router.Use(cors.Handler)
	router.Use(
		middleware.Logger,
		appmiddleware.Recoverer,
	)

	// Initialize the version 1 routes of the API
	router.Route("/v1", v1.Init)

	// TODO: register with auctioneer

	trace.Log.Infof("Starting bidder %s on port :%d with delay %d\n",
		*name, *port, *delay)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
