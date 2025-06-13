package main

import (
	"github.com/IvanDrf/currency-aggregator/internal/handlers"
)

func main() {
	server := handlers.InitServer()
	server.RegisterRoutes()

	server.Start()
}
