package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
)

type UserRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

type UserRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func BindUserRoutes(r *gin.Engine, us service.UserServicer) {
	userRouter := r.Group("/api/v1/users")

	createUser(userRouter, us)
	getUser(userRouter, us)
	getUsers(userRouter, us)
	editUser(userRouter, us)
}

func createUser(r *gin.RouterGroup, us service.UserServicer) {
	createUserHandler := func(c *gin.Context) {
		var req UserRequestBody

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		user := us.CreateUser(req.Name, req.Email, req.Password)
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
	r.POST("", createUserHandler)
}

func getUser(r *gin.RouterGroup, us service.UserServicer) {
	getUserHandler := func(c *gin.Context) {
		var req UserRequestParam

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
		var reqId UserRequestParam
		var reqBody UserRequestBody

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
