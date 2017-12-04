package router

import (
	"github.com/gin-gonic/gin"
	. "gin-mysql-restful/router/group"
	. "gin-mysql-restful/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	InitPolicyRouter(router)

	return router
}
