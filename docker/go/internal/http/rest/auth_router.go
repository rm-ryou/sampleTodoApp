package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
)

type SignInParam struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func BindAuthRoutes(r *gin.Engine, as service.AuthServicer) {
	authRouter := r.Group("/api/v1/auth")

	usersSignIn(authRouter, as)
}

func usersSignIn(r *gin.RouterGroup, as service.AuthServicer) {
	signInHandler := func(c *gin.Context) {
		var req SignInParam

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		res, err := as.SignIn(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
	r.POST("/users/sign_in", signInHandler)
}
