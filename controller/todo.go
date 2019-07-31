package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mw "github.com/structe/go-meeting/middleware"
	"github.com/structe/go-meeting/model"
	"github.com/structe/go-meeting/utils"
)

type Postdate struct {
	Userid int    `json:"userid"`
	Date   string `json:"date"`
}

func Newtodo(c *gin.Context) {
	var (
		todo model.Todo
		res  = gin.H{}
		err  error
	)

	err = mw.CheckToken(c)
	if err != nil {
		fmt.Println(err)
		utils.SendBadResponse(c, "token invalid")
		return
	}

	if err = c.ShouldBindJSON(&todo); err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}

	err = todo.Insert()
	if err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}

	res["messgae"] = "succeed"
	c.JSON(http.StatusOK, res)
}

func Updatetodo(c *gin.Context) {
	var (
		todo model.Todo
		res  = gin.H{}
		err  error
	)

	err = mw.CheckToken(c)
	if err != nil {
		println(err)
		utils.SendBadResponse(c, "token invalid")
		return
	}

	if err = c.ShouldBindJSON(&todo); err != nil {
		//println("internal error")
		utils.SendBadResponse(c, "internal error")
		return
	}

	err = todo.Update()
	if err != nil {
		utils.SendBadResponse(c, "update error")
		return
	}

	res["messgae"] = "succeed"
	c.JSON(http.StatusOK, res)
}

func Deletetodo(c *gin.Context) {
	var (
		todo model.Todo
		res  = gin.H{}
		err  error
	)

	err = mw.CheckToken(c)
	if err != nil {
		utils.SendBadResponse(c, "token invalid")
		return
	}
	if err = c.ShouldBindJSON(&todo); err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}
	err = model.DeleteTodo(todo.ID)
	if err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}

	res["messgae"] = "succeed"
	c.JSON(http.StatusOK, res)
}

func ListTodoByDate(c *gin.Context) {
	var (
		todos []*model.Todo
		data  Postdate
		res   = gin.H{}
		err   error
	)

	err = mw.CheckToken(c)
	if err != nil {
		utils.SendBadResponse(c, "token invalid")
		return
	}

	if err = c.ShouldBindJSON(&data); err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}
	todos, err = model.GetTodoByDate(data.Userid, data.Date)
	if err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}

	res["messgae"] = "succeed"
	res["data"] = todos
	c.JSON(http.StatusOK, res)
}

func ListTodoByUser(c *gin.Context) {
	var (
		todos []*model.Todo
		data  Postdate
		res   = gin.H{}
		err   error
	)

	err = mw.CheckToken(c)
	if err != nil {
		utils.SendBadResponse(c, "token invalid")
		return
	}

	if err = c.ShouldBindJSON(&data); err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}
	todos, err = model.GetAll(data.Userid)
	if err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}

	res["messgae"] = "succeed"
	res["data"] = todos
	c.JSON(http.StatusOK, res)
}
