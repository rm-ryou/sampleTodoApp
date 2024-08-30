package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
)

func BindUserRoutes(r *gin.Engine, us service.UserServicer) {
	userRouter := r.Group("/api/v1/users")

	getUsers(userRouter, us)
}

func getUsers(r *gin.RouterGroup, us service.UserServicer) {
	getUsersHandler := func(c *gin.Context) {
		users := us.GetUsers()
		c.JSON(200, gin.H{"data": users})
	}
	r.GET("", getUsersHandler)
}
