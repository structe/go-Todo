package router

import (
	"github.com/gin-gonic/gin"
	"github.com/structe/go-meeting/controller"
)

func Route(router *gin.Engine) {
	api := router.Group("/meet/api")

	{
		api.POST("/signin", controller.SignIn)
		api.POST("/signup", controller.SignUp)
		api.POST("/newtodo", controller.Newtodo)
		api.POST("/updatetodo", controller.Updatetodo)
		api.POST("/deletetodo", controller.Deletetodo)
		api.POST("/listtodo", controller.ListTodoByDate)
		api.POST("/listall", controller.ListTodoByUser)
	}
}
