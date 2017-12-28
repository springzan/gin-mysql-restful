package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "gin-mysql-restful/boot"
)

func IndexApi(c *gin.Context) {
	accept:=SignNoUser(c)
	if !accept {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "ok",
	})
}
