package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/http/rest"
	"github.com/rm-ryou/sampleTodoApp/pkg/auth"
)

var cacheApi *Api

type Api struct {
	Router *gin.Engine
}

func GetApi() *Api {
	return cacheApi
}

func Initialize() {
	signingKey := os.Getenv("SIGNING_KEY")
	if signingKey == "" {
		log.Fatalln("signingKey is not defined")
	}
	auth.InitializeSigningKey(signingKey)

	r := rest.NewRouter()

	rest.BindRoutes(r)

	cacheApi = &Api{r}
}

func Host() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("Port is not configured")
	}

	return fmt.Sprintf(":%s", port)
}
