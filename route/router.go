package route

import (
	"code.lhc.org/we-blog/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// user
	router.GET("/", api.GetIndex)
	router.GET("/user/login", api.Login)
	router.POST("/user/register", api.Register)
	//router.GET("/user/info")
	// comment

	// article

	return router
}
