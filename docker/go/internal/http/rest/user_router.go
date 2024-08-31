package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
	"github.com/rm-ryou/sampleTodoApp/pkg/middleware"
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
	userRouter.Use(middleware.AuthMiddleware())

	getUser(userRouter, us)
	getUsers(userRouter, us)
	editUser(userRouter, us)
	deleteUser(userRouter, us)
}

func getUser(r *gin.RouterGroup, us service.UserServicer) {
	getUserHandler := func(c *gin.Context) {
		var req UserRequestParam

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		user, err := us.GetUser(req.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
	r.GET("/:id", getUserHandler)
}

func getUsers(r *gin.RouterGroup, us service.UserServicer) {
	getUsersHandler := func(c *gin.Context) {
		users, err := us.GetUsers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		}

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

		user, err := us.EditUser(reqId.ID, reqBody.Name, reqBody.Email, reqBody.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
	r.PATCH("/:id", editUserHandler)
}

func deleteUser(r *gin.RouterGroup, us service.UserServicer) {
	deleteUserHandler := func(c *gin.Context) {
		var reqId UserRequestParam

		if err := c.ShouldBindUri(&reqId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		if err := us.DeleteUser(reqId.ID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	}
	r.DELETE("/:id", deleteUserHandler)
}
