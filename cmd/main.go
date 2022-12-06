package main

import (
	"go-rest-template/internal/router"
	"log"
)

func main() {
	r := router.New()

	if err := r.Run("localhost:8000"); err != nil {
		log.Fatalln(err)
	}
}
