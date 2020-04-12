package main

import (
	"fmt"
	"os"

	"to_do_list/db"
	todo "to_do_list/module/todo/routes"

	"github.com/jinzhu/gorm"
)

var (
	err error
)

func initMysql() *gorm.DB {
	DB, err := db.OpenConnectionMysql()
	if err != nil {
		fmt.Println("[error]", err.Error())
		os.Exit(1)
	}
	fmt.Print("Mysql connection initialized")
	return DB
}

func main() {
	DB := initMysql()
	// db.AutoMigrateMysql(DB, &models.ToDo{}, &models.Users{})
	r := todo.SetupRouter(DB)
	//running server
	r.Run()
}
