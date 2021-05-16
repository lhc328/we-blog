package api

import (
	"code.lhc.org/we-blog/common"
	"code.lhc.org/we-blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	resp := make(map[string]interface{}, 0)
	userSrv := service.UserService{}
	user, token, err := userSrv.LoginByPassword(c)
	resp["data"] = user.Id
	if err == nil {
		c.Header("we-token", token)
		common.CommonSuccess(c, resp)
	} else {
		fmt.Printf("[login error] err is %s", err)
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

func UserInfo(c *gin.Context) {
	resp := make(map[string]interface{}, 0)
	userSrv := service.UserService{}
	user, err := userSrv.GetUserInfo(c)
	resp["data"] = user
	if err == nil {
		common.CommonSuccess(c, resp)
	} else {
		fmt.Printf("[user error] err is %s", err)
		common.CommonFail(c, resp, "获取失败")
	}
}
