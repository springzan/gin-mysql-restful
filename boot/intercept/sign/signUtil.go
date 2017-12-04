package sign

import (
	"fmt"
	"crypto/md5"
	"io"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

const AUTH_KEY  = "EBT-Authorization"

type AuthorizationInfo struct {
	DeviceId string `json:"deviceId" form:"deviceId"`
	SignTimestamp int64 `json:"signTimestamp" form:"signTimestamp"`
	SignVersion string `json:"signVersion" form:"signVersion"`
	Sign string `json:"sign" form:"sign"`
}

func ParseAuthInfo(c *gin.Context) AuthorizationInfo {
	authStr:= c.GetHeader(AUTH_KEY)
	var authorizationInfo AuthorizationInfo
	if err := json.Unmarshal([]byte(authStr), &authorizationInfo); err == nil {
		return authorizationInfo
	}else {
		fmt.Println("error==============json str 转struct AuthorizationInfo ==")
		fmt.Println(authStr)
	}
	return authorizationInfo
}

func CheckSign(authorizationInfo AuthorizationInfo) bool{
	signVersion:=authorizationInfo.SignVersion
	originKey:=getOriginKeyBySignVersion(signVersion)
	if originKey==""{
		return false
	}
	sign:=authorizationInfo.Sign
	deviceId:= authorizationInfo.DeviceId
	signTimestamp:= authorizationInfo.SignTimestamp
	requestKey:= md5Sum(deviceId + originKey)
	serverSign := md5Sum(deviceId + strconv.FormatInt(signTimestamp,10)   + requestKey)
	return serverSign==sign
}

func getOriginKeyBySignVersion(signVersion string) string {
	var SignKeyMap = map[string]string{
		"v1.1":"ios-zyj-v1.1",
		"zyj-iOS-1.4": "4S44-SS-4446544",
		"zyj-Android-1.4": "8-rind8rid-n7rd87-i",
		"zyj-iOS-2.0.0": "4S44-SS-4446544",
		"zyj-Android-2.0.0": "8-rind8rid-n7rd87-i",
		"xhybt-iOS-1.0.0": "hbb-9.--hb-b-.b-bbb",
		"xhybt-Android-1.0.0": "AAA-7n.7.dd.40.4dn0-.A",
		"zyj-sxbk-0.0.1": "j.8s8.sss83ss884ss",
	}
	return SignKeyMap[signVersion]
}

func md5Sum(key string) string {
	w := md5.New()
	io.WriteString(w, key)   //将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))  //w.Sum(nil)将w的hash转成[]byte格式
	return md5str2
}

const SIGN_EXPIRE_SECONDS = "-10s"

///验证签名是否过期; 如果过期, 返回true
func CheckTimeout(authorizationInfo AuthorizationInfo) bool {
	signTimestamp:= authorizationInfo.SignTimestamp
	fmt.Println(signTimestamp)
	if signTimestamp==0{
		return false
	}
	//time.Sleep(2*1e9)
	now:=time.Now()
	loc, err := time.LoadLocation("Local")
	if err!=nil{
		return false
	}
	//nowInt64Str:=strconv.FormatInt(now.In(loc).UnixNano(), 10)[:13]
	//fmt.Println(nowInt64Str)
	s, _ := time.ParseDuration(SIGN_EXPIRE_SECONDS)
	tenSecondAgo:= now.Add(s)
	tenSecondAgoStr:=strconv.FormatInt(tenSecondAgo.In(loc).UnixNano(), 10)[:13]
	tenSecondAgoInt64, err := strconv.ParseInt(tenSecondAgoStr, 10, 64)
	fmt.Println(tenSecondAgoInt64)
	return signTimestamp < tenSecondAgoInt64
}