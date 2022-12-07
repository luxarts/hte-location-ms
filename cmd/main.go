package main

import (
	"go-rest-template/internal/defines"
	"go-rest-template/internal/router"
	"log"
	"os"
)

func main() {
	r := router.New()

	if err := r.Run(os.Getenv(defines.EnvAPILocation)); err != nil {
		log.Fatalln(err)
	}
}
