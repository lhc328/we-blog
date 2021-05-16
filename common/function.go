package common

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CommonFail(c *gin.Context, backData map[string]interface{}, message string) {
	backData["code"] = 5000
	backData["message"] = message
	c.JSON(200, backData)
}

func CommonSuccess(c *gin.Context, backData map[string]interface{}) {
	backData["code"] = 2000
	backData["message"] = "success"
	c.JSON(200, backData)
}

func GetAuth2Session(code string) (map[string]interface{}, int64) {
	return nil, 0
}

func Md5(value string) string {
	data := md5.Sum([]byte(value))
	return fmt.Sprintf("%x", data)
}
