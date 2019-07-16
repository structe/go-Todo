package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 错误请求
func SendBadResponse(c *gin.Context, err string) {
	c.JSON(http.StatusOK, gin.H{
		"msg": err,
	})
	// 中止请求链
	c.Abort()
}

func SHA256Str(src string) string {
	h := sha256.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}
