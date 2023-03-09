package main

import (
	"hte-location-ms/internal/metrics"
	"hte-location-ms/internal/router"
	"log"
)

func main() {
	metrics.StartServer()

	r := router.New()

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
