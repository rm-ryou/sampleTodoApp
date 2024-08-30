package main

import (
	"log"

	"github.com/rm-ryou/sampleTodoApp/internal/http/rest"
)

func main() {
	router := rest.NewRouter()
	if err := router.Run("3100"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
