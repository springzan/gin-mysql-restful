package policydb

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"gin-mysql-restful/config"
)

var PolicyDB *sql.DB

const dbCfNode  = "policy"

func init(){
	dbConfig:=config.GetDBConfig(dbCfNode)
	var err error
	PolicyDB, err=sql.Open(dbConfig.Dialect, dbConfig.URL)
	if err != nil{
		log.Fatal(err.Error())
	}
	PolicyDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	PolicyDB.SetMaxIdleConns(dbConfig.MaxIdleConns)

	err=PolicyDB.Ping()
	if err != nil{
		log.Fatal(err.Error())
	}
}
