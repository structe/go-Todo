package router

import (
	"github.com/gin-gonic/gin"
	user "github.com/structe/go-meeting/controller"
)

func Route(router *gin.Engine) {
	api := router.Group("/api")

	{
		api.POST("/signin", user.SignIn)
		api.POST("/signup", user.SignUp)
	}
}
