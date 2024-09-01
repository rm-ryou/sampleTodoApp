package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
	"github.com/rm-ryou/sampleTodoApp/internal/storage/mysql/repository"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	return r
}

func BindRoutes(r *gin.Engine) {
	ur := repository.NewUserRepository()
	us := service.NewUserService()
	as := service.NewAuthService(ur)

	BindUserRoutes(r, us)
	BindAuthRoutes(r, as)
}
