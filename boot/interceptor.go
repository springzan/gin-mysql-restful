package boot

/*import (
	"github.com/gin-gonic/gin"
	. "gin-mysql-restful/boot/intercept"
	"net/http"
)

func NoSignNoUser(c *gin.Context) bool {
	return true
}

func SignButNoUser(c *gin.Context) bool{
	s,err:=SignPreHandle(c)
	if !s {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return false
	}
	return true
}

func NoSignButUser(c *gin.Context) bool {
	u,err:=UserIdentityPreHandle(c)
	if !u {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return false
	}
	return true
}

func SignAndUser(c *gin.Context)  bool{
	s,sErr:=SignPreHandle(c)
	if !s {
		c.JSON(http.StatusOK, gin.H{
			"error": sErr,
		})
		return false
	}
	u,uErr:=UserIdentityPreHandle(c)
	if !u {
		c.JSON(http.StatusOK, gin.H{
			"error": uErr,
		})
		return false
	}
	return true
}*/

