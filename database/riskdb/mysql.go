package database

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"gin-mysql-restful/config"
)

var RiskDB *sql.DB

const dbCfNode  = "risk"

func init(){
	dbConfig:=config.GetDBConfig(dbCfNode)
	var err error
	RiskDB, err=sql.Open(dbConfig.Dialect, dbConfig.URL)
	if err != nil{
		log.Fatal(err.Error())
	}
	RiskDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	RiskDB.SetMaxIdleConns(dbConfig.MaxIdleConns)

	err=RiskDB.Ping()
	if err != nil{
		log.Fatal(err.Error())
	}
}
