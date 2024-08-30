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
	router = NewRouter()

	userService := service_mock.NewUserServiceMock()
	BindUserRoutes(router, userService)

	m.Run()
}
