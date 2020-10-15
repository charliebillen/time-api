package main

import (
	"net/http"
	"time"

	api "github.com/charliebillen/time-api"
)

const oneSecond = 1 * time.Second

func main() {
	timeServer := &api.Server{}

	httpServer := http.Server{
		Addr:         ":8000",
		ReadTimeout:  oneSecond,
		WriteTimeout: oneSecond,
		IdleTimeout:  oneSecond,
		Handler:      timeServer,
	}

	httpServer.ListenAndServe()
}