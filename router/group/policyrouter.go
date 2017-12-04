package group

import (
	"github.com/gin-gonic/gin"
	. "gin-mysql-restful/apis/policy"
	)

func InitPolicyRouter(router *gin.Engine)  {
	policyRouter:=router.Group("/policy")
	{
		policyRouter.GET("/:policyId", GetPolicyApi)
	}
}