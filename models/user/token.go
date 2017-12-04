package user

import (
	. "gin-mysql-restful/database/userdb"
)

type Token struct {
	UserId int `json:"userId" form:"policyId"`
	Token string `json:"token" form:"token"`
}

func GetOneByToken(token string) (t Token,err error) {
	row := UserDB.QueryRow("select user_id,token from token where token =?", token)
	if row == nil {
		return
	}
	err = row.Scan(&t.UserId, &t.Token)
	if err != nil {
		return
	}
	return
}
