package main

import (
	"fmt"

	"to_do_list/db"
	"to_do_list/module/todo/models"
	"to_do_list/module/todo/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	db.DB, err = gorm.Open("mysql", db.DbURL(db.BuildDbConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}
	defer db.DB.Close()
	db.DB.AutoMigrate(&models.ToDo{})
	r := routes.SetupRouter()
	//running server
	r.Run()
}
