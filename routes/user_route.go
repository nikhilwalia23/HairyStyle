package routes

import (
	"my-go-server/controllers"
	"my-go-server/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("/register", controllers.RegisterUser)
		users.POST("/login", controllers.Login)
		users.POST("/protected", utils.JWTAuthMiddleware, controllers.ProtectedTesting)

	}
}
