package main

import (
	"log"

	"github.com/rm-ryou/sampleTodoApp/internal/config"
)

func main() {
	config.Initialize()

	err := config.GetApi().Router.Run(config.Host())
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
