package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/structe/go-meeting/config"
	"github.com/structe/go-meeting/model"
	"github.com/structe/go-meeting/router"
)

func main() {
	fmt.Println("gin version: ", gin.Version)

	config.InitConfig()

	//gin.SetMode(gin.ReleaseMode)

	// Creates a router without any middleware by default
	app := gin.Default()

	db, err := model.InitMysql()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	router.Route(app)

	app.Run(":" + config.ServerPort)
}
