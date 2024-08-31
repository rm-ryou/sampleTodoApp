package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	return r
}

func BindRoutes(r *gin.Engine) {
	us := service.NewUserService()
	as := service.NewAuthService()

	BindUserRoutes(r, us)
	BindAuthRoutes(r, as)
}
