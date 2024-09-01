package rest

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
	"github.com/rm-ryou/sampleTodoApp/internal/storage/mysql/repository"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	return r
}

func BindRoutes(r *gin.Engine, db *sql.DB) {
	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)
	as := service.NewAuthService(ur)

	BindUserRoutes(r, us)
	BindAuthRoutes(r, as)
}
