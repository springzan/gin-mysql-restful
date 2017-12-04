package config

import (
	"io/ioutil"
	"fmt"
	"os"
	"regexp"
	"encoding/json"
	"gin-mysql-restful/utils"
)

var jsonData map[string]interface{}

func initJSON() {
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg       := regexp.MustCompile(`/\*.*\*/`)

	configStr  = reg.ReplaceAllString(configStr, "")
	bytes      = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

type DBConfig struct {
	Dialect       string
	Database      string
	User          string
	Password      string
	Host          string
	Port          int
	Charset       string
	URL           string
	MaxIdleConns  int
	MaxOpenConns  int
}

// DBConfig 数据库相关配置
var dbConfig DBConfig

var database map[string]interface{}

func initDB(dbName string) DBConfig{
	utils.ParseStructByJSON(&dbConfig, database[dbName].(map[string]interface{}))
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Charset)
	dbConfig.URL = url
	return dbConfig
}

func init()  {
	initJSON()
	database=jsonData["database"].(map[string]interface{})
}

func GetDBConfig(dbName string) DBConfig{
	return initDB(dbName)
}