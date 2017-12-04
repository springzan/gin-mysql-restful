package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"gin-mysql-restful/config"
)

var UserDB *sql.DB

const dbName  = "user"

func init(){
	dbConfig:=config.GetDBConfig(dbName)
	var err error
	UserDB, err=sql.Open(dbConfig.Dialect, dbConfig.URL)
	if err != nil{
		log.Fatal(err.Error())
	}
	UserDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	UserDB.SetMaxIdleConns(dbConfig.MaxIdleConns)

	err=UserDB.Ping()
	if err != nil{
		log.Fatal(err.Error())
	}
}
