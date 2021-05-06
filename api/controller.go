package api

import (
	"code.lhc.org/we-blog/common"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	code := c.DefaultQuery("code", "0")
	var backData map[string]interface{}
	if code == "0" {
		common.CommonFail(c, backData, "code值缺少")
	} else {
		backdata, status :=common.GetAuth2Session(code)
		if status == 0 {
			common.CommonFail(c, backdata, "code 验证失败")
		} else {
			common.CommonSuccess(c, backdata)
		}
	}
}
