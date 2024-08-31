package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
)

type GetUserRequest struct {
	ID int `uri:"id" binding:"required"`
}

type EditUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func BindUserRoutes(r *gin.Engine, us service.UserServicer) {
	userRouter := r.Group("/api/v1/users")

	getUser(userRouter, us)
	getUsers(userRouter, us)
	editUser(userRouter, us)
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

func editUser(r *gin.RouterGroup, us service.UserServicer) {
	editUserHandler := func(c *gin.Context) {
		var reqId GetUserRequest
		var reqBody EditUserRequest

		if err := c.ShouldBindUri(&reqId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		user := us.EditUser(reqId.ID, reqBody.Name, reqBody.Email, reqBody.Password)
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
	r.PATCH("/:id", editUserHandler)
}
