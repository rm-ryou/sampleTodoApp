package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rm-ryou/sampleTodoApp/internal/service"
)

type SignInParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func BindAuthRoutes(r *gin.Engine, as service.AuthServicer) {
	authRouter := r.Group("/api/v1/auth")

	authRouter.POST("/users/sign_in", signInHandler(as, false))
	authRouter.POST("/admins/sign_in", signInHandler(as, true))
}

func signInHandler(as service.AuthServicer, isAdminResource bool) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req SignInParam

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		res, err := as.SignIn(req.Email, req.Password, isAdminResource)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
