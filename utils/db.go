package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbHandle gorm.DB

func ConnDB(conf *ServerConf) gorm.DB {
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DataBase.User, conf.DataBase.Password, conf.Server.Ip, conf.DataBase.Port, conf.DataBase.Dbname)
	log.Printf("db connect param is %s\n", mysqlDsn)
	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if err != nil {
		log.Printf("connect to database error: %s\n", err.Error())
		os.Exit(1)
	}
	return *db
}
func GetDBHandle() gorm.DB {
	return dbHandle
}
