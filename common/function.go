package common

import "github.com/gin-gonic/gin"

func CommonFail(c *gin.Context, backData map[string]interface{}, message string) {
	c.JSON(500, gin.H{"message": message})
}

func CommonSuccess(c *gin.Context, backData map[string]interface{}) {
	c.JSON(200, gin.H{"message": "success"})
}

func GetAuth2Session(code string) (map[string]interface{}, int64) {
	return nil, 0
}
