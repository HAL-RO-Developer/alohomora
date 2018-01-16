package model

import (
	"io/ioutil"
	"os"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

var db = NewDBConn()

func NewDBConn() *gorm.DB {
	driver, resource := "mysql", "root:password@tcp(127.0.0.1:3306)/lavender?parseTime=true&collation=utf8_general_ci&interpolateParams=true&loc=Local"
	if os.Getenv("lavender") == "develop" {
		driver, resource = GetDBConfig("dbconfig.yml", "develop")
	}
	if os.Getenv("lavender") == "production" {
		driver, resource = GetDBConfig("dbconfig.yml", "development")
	}
	db, err := gorm.Open(driver, resource)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB: "+resource)
	return db
}

func GetDBConn() *gorm.DB {
	return db
}

func GetDBConfig(configPath string, dbname string) (string, string) {
	var buf, err = ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
	driver := m[dbname].(map[interface{}]interface{})["dialect"].(string)
	source := m[dbname].(map[interface{}]interface{})["datasource"].(string)
	return driver, source
}
