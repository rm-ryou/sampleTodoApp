package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/pkg/auth"
	"github.com/rm-ryou/sampleTodoApp/pkg/utils"
	service_mock "github.com/rm-ryou/sampleTodoApp/test/mock/service"
)

var (
	router  *gin.Engine
	port    = os.Getenv("PORT")
	baseURL = fmt.Sprintf("http://localhost:%s", port)
)

func setHeader(userId int, r *http.Request) error {
	token, err := auth.GenerateToken(userId, utils.RealTime{})
	if err != nil {
		return err
	}

	r.Header.Set("Authorization", "Bearer "+token)

	return nil
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	router = NewRouter()

	userService := service_mock.NewUserServiceMock()
	authService := service_mock.NewAuthServiceMock()
	BindUserRoutes(router, userService)
	BindAuthRoutes(router, authService)

	m.Run()
}
