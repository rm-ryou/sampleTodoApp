package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/http/rest"
	"github.com/rm-ryou/sampleTodoApp/internal/storage/mysql"
	"github.com/rm-ryou/sampleTodoApp/pkg/auth"
)

var cacheApi *Api

type Api struct {
	Router *gin.Engine
	DB     *sql.DB
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
	db, err := mysql.SetUpDB(getDBConfig())
	if err != nil {
		log.Fatal(err.Error())
	}

	rest.BindRoutes(r, db)

	cacheApi = &Api{r, db}
}

func Host() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("Port is not configured")
	}

	return fmt.Sprintf(":%s", port)
}

func getDBConfig() (string, string) {
	driver := os.Getenv("DB_DRIVER")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return driver, dsn
}
