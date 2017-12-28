package boot

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	."gin-mysql-restful/boot/intercept/sign"
	."gin-mysql-restful/models/user"
)

type RequestCheck interface {
	PreHandle(c *gin.Context) bool
}

type SignError struct{
	Code string `json:"code"`
	Message string `json:"message"`
	Type string `json:"type"`
}

type UserIdentityError struct{
	Code string `json:"code"`
	Message string `json:"message"`
	Type string `json:"type"`
}


func (err SignError) PreHandle(c *gin.Context) bool{
	author:= ParseAuthInfo(c)
	if author.Sign=="" {
		fmt.Println("签名信息缺少必要参数")
		error := SignError{ Code:"100101",Message:"签名信息缺少必要参数",Type:"signParamIllegal"}
		c.JSON(http.StatusOK, gin.H{
			"error": error,
		})
		return false
	}
	if !author.CheckSign(){
		fmt.Println("签名不正确")
		error := SignError{ Code:"100102",Message:"签名不正确",Type:"signIllegal"}
		c.JSON(http.StatusOK, gin.H{
			"error": error,
		})
		return false
	}
	if author.CheckTimeout(){
		fmt.Println("签名超时，签名时间跟当前时间间隔太长，需要重签")
		error := SignError{ Code:"100103",Message:"签名超时，签名时间跟当前时间间隔太长，需要重签",Type:"signTimeout"}
		c.JSON(http.StatusOK, gin.H{
			"error": error,
		})
		return false
	}
	return true
}

const TOKEN  ="Token"

func (err UserIdentityError)PreHandle(c *gin.Context) bool {
	token:= c.GetHeader(TOKEN)
	if token==""{
		error:=UserIdentityError{"100201","Token为空","invalidToken"}
		c.JSON(http.StatusOK, gin.H{
			"error": error,
		})
		return false
	}
	userToken, _ :=GetOneByToken(token)
	if userToken.UserId<1{
		error:=UserIdentityError{"100202","无效的Token","invalidToken"}
		c.JSON(http.StatusOK, gin.H{
			"error": error,
		})
		return false
	}
	return true
}

func preHandles(checks []RequestCheck, c *gin.Context) bool {
	for _, v := range checks {
		return v.PreHandle(c)
	}
	return true
}

func SignUser(c *gin.Context) bool {
	var s SignError
	var u UserIdentityError
	checks:=[]RequestCheck{s,u}
	return preHandles(checks,c)
}

func SignNoUser(c *gin.Context) bool{
	var s SignError
	checks:=[]RequestCheck{s}
	return preHandles(checks,c)
}