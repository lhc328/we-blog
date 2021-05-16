package route

import (
	"code.lhc.org/we-blog/api"
	"code.lhc.org/we-blog/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// user
	router.GET("/", middleware.JWTAuth(), api.GetIndex)
	router.POST("/user/login", api.Login)
	router.POST("/user/register", api.Register)
	router.GET("/user/info", middleware.JWTAuth(), api.UserInfo)
	// comment

	// article

	return router
}
