package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

//DBconfig :nodoc:
type DBconfig struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

//ViperConfigVariable read by key viper
func ViperConfigVariable(key string) string {

	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatal("Invalid type assertion")
	}
	return value
}

//BuildDbConfig :nodoc:
func BuildDbConfig() *DBconfig {
	DBconfig := DBconfig{
		Host:     ViperConfigVariable("db_host"),
		Port:     ViperConfigVariable("db_port"),
		User:     ViperConfigVariable("db_user"),
		DbName:   ViperConfigVariable("db_name"),
		Password: ViperConfigVariable("db_password"),
	}
	return &DBconfig
}

//DbURL :nodoc:
func DbURL(DBconfig *DBconfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DBconfig.User,
		DBconfig.Password,
		DBconfig.Host,
		DBconfig.Port,
		DBconfig.DbName,
	)

}
