package rest

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	service_mock "github.com/rm-ryou/sampleTodoApp/test/mock/service"
)

var (
	router  *gin.Engine
	port    = os.Getenv("PORT")
	baseURL = fmt.Sprintf("http://localhost:%s", port)
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	router = NewRouter()

	userService := service_mock.NewUserServiceMock()
	authService := service_mock.NewAuthServiceMock()
	BindUserRoutes(router, userService)
	BindAuthRoutes(router, authService)

	m.Run()
}
