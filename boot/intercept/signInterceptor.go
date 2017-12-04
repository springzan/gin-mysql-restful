package intercept

import (
	. "gin-mysql-restful/boot/intercept/sign"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SignError struct{
	Code string `json:"code"`
	Message string `json:"message"`
	Type string `json:"type"`
}

func SignPreHandle(c *gin.Context) (bool,SignError) {
	author:= ParseAuthInfo(c)
	if author.Sign=="" {
		fmt.Println("签名信息缺少必要参数")
		error := SignError{ Code:"100101",Message:"签名信息缺少必要参数",Type:"signParamIllegal"}
		return false,error
	}
	if !CheckSign(author){
		fmt.Println("签名不正确")
		error := SignError{ Code:"100102",Message:"签名不正确",Type:"signIllegal"}
		return false,error
	}

	if CheckTimeout(author){
		fmt.Println("签名超时，签名时间跟当前时间间隔太长，需要重签")
		error := SignError{ Code:"100103",Message:"签名超时，签名时间跟当前时间间隔太长，需要重签",Type:"signTimeout"}
		return false,error
	}
	error := SignError{}
	return true,error
}
