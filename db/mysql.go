package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DB as a driver gorm
var DB *gorm.DB

//OpenConnectionMysql open connection to mysql
func OpenConnectionMysql() (*gorm.DB, error) {
	DB, err := gorm.Open("mysql", MysqlConnURL(BuildDbConfig()))
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func AutoMigrateMysql(db *gorm.DB, models ...interface{}) error {
	if models == nil {
		return nil
	}
	db.AutoMigrate(models)
	return nil
}
