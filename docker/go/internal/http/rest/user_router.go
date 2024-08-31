package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
)

type GetUserRequest struct {
	ID int `uri:"id" binding:"required"`
}

func BindUserRoutes(r *gin.Engine, us service.UserServicer) {
	userRouter := r.Group("/api/v1/users")

	getUser(userRouter, us)
	getUsers(userRouter, us)
}

func getUser(r *gin.RouterGroup, us service.UserServicer) {
	getUserHandler := func(c *gin.Context) {
		var req GetUserRequest

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		user := us.GetUser(req.ID)
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
	r.GET("/:id", getUserHandler)
}

func getUsers(r *gin.RouterGroup, us service.UserServicer) {
	getUsersHandler := func(c *gin.Context) {
		users := us.GetUsers()
		c.JSON(http.StatusOK, gin.H{"data": users})
	}
	r.GET("", getUsersHandler)
}
