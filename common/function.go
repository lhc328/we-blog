package common

import "github.com/gin-gonic/gin"

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
