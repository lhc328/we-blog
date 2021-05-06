package route

import (
	"code.lhc.org/we-blog/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", api.GetIndex)

	return router
}