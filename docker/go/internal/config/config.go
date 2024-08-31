package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/http/rest"
)

var cacheApi *Api

type Api struct {
	Router *gin.Engine
}

func GetApi() *Api {
	return cacheApi
}

func Initialize() {
	r := rest.NewRouter()

	rest.BindRoutes(r)

	cacheApi = &Api{r}
}

func Host() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("Port is not configured")
	}

	return fmt.Sprintf(":%s", port)
}
