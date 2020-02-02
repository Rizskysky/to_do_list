package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBconfig struct {
	Host     string
	Port     int
	User     string
	DbName   string
	Password string
}

func BuildDbConfig() *DBconfig {
	DBconfig := DBconfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "rizsky",
		DbName:   "todo",
		Password: "12345678",
	}
	return &DBconfig
}

//function
func DbURL(DBconfig *DBconfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		DBconfig.User,
		DBconfig.Password,
		DBconfig.Host,
		DBconfig.Port,
		DBconfig.DbName,
	)

}
