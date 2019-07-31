package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/structe/go-meeting/config"
	"github.com/structe/go-meeting/utils"
)

type CustomClaims struct {
	UserName string `json:"username"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}

//创建jwttoken
func CreateJwtToken(c *gin.Context, ID int, UserName string) {
	// Create the Claims
	claims := CustomClaims{
		UserName,
		ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "meet",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.TokenSecret))
	if err != nil {
		utils.SendBadResponse(c, "internal error")
		return
	}
	SetCookie(c, "token", tokenString)
}

//解析jwt
func ParseJwtToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenSecret), nil
	})
	_, ok := token.Claims.(*CustomClaims)
	return ok && token.Valid && err == nil
}

//设置cookie
func SetCookie(c *gin.Context, name string, value string) {
	c.SetCookie(name, value, config.CooKieMaxAge*60*60, "/", "", false, true)
}

func CheckToken(c *gin.Context) error {
	token, err := c.Cookie("token")
	if err != nil {
		return err
	}
	flag := ParseJwtToken(token)
	if flag != true {
		return err
	}
	return err
}
