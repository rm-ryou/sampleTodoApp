package middleware

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/pkg/auth"
)

var (
	router     *gin.Engine
	signingKey = "test"
	baseURL    = fmt.Sprintf("http://localhost:%s", os.Getenv("PORT"))
)

func bindAuthRoutes() {
	authRouter := router.Group("/auth")
	authRouter.Use(AuthMiddleware())

	testAuthHandler := func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Access granted"})
	}
	authRouter.GET("", testAuthHandler)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	auth.InitializeSigningKey(signingKey)
	router = gin.New()
	bindAuthRoutes()

	m.Run()
}
