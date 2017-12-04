package intercept

import (
	"github.com/gin-gonic/gin"
	. "gin-mysql-restful/models/user"
)

type UserIdentityError struct{
	Code string `json:"code"`
	Message string `json:"message"`
	Type string `json:"type"`
}

const TOKEN  ="Token"

func UserIdentityPreHandle(c *gin.Context) (bool,UserIdentityError) {
	token:= c.GetHeader(TOKEN)
	if token==""{
		err:=UserIdentityError{"100201","Token为空","invalidToken"}
		return false,err
	}
	userToken, _ :=GetOneByToken(token)
	if userToken.UserId<1{
		err:=UserIdentityError{"100202","无效的Token","invalidToken"}
		return false,err
	}
	err:=UserIdentityError{}
	return true,err
}