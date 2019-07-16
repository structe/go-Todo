package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mw "github.com/structe/go-meeting/middleware"
	"github.com/structe/go-meeting/model"
	"github.com/structe/go-meeting/utils"
)

type SignInData struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func SignUp(c *gin.Context) {
	var (
		user model.User
		res  = gin.H{}
		err  error
	)

	if err = c.BindJSON(&user); err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}
	if len(user.Email) == 0 || len(user.Password) == 0 {
		utils.SendBadResponse(c, "email or password cannot be null")
		return
	}
	user.Password = utils.SHA256Str(user.Email + user.Password)
	err = user.Insert()
	if err != nil {
		utils.SendBadResponse(c, "email exists")
		return
	}
	res["messgae"] = "succeed"
	c.JSON(http.StatusOK, res)
}

func SignIn(c *gin.Context) {
	var (
		user *model.User
		res  = gin.H{}
		err  error
	)
	if err = c.BindJSON(&user); err != nil {
		utils.SendBadResponse(c, "internal error")
		c.JSON(http.StatusOK, res)
		return
	}
	email := user.Email
	password := user.Password
	if email == "" || password == "" {
		utils.SendBadResponse(c, "username or password cannot be null")
		return
	}

	user, err = model.GetUserByEmail(email)
	if err != nil || user.Password != utils.SHA256Str(email+password) {
		utils.SendBadResponse(c, "invalid username or password")
		return
	}
	mw.CreateJwtToken(c, user.ID, user.Email)
	res["message"] = "succeed"
	c.JSON(http.StatusOK, res)
}
