package api

import (
	"code.lhc.org/we-blog/common"
	"code.lhc.org/we-blog/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	resp := make(map[string]interface{}, 0)
	userSrv := service.UserService{}
	user, err := userSrv.LoginByPassword(c)
	resp["data"] = user
	if err == nil {
		common.CommonSuccess(c, resp)
	} else {
		common.CommonFail(c, resp, "登录失败")
	}
}

func Register(c *gin.Context) {
	resp := make(map[string]interface{}, 0)
	userSrv := service.UserService{}
	id, err := userSrv.RegisterByPassword(c)
	resp["data"] = map[string]int64{"id": id}
	if err == nil {
		common.CommonSuccess(c, resp)
	} else {
		common.CommonFail(c, resp, "登录失败")
	}
}
