package main

import (
	"fmt"

	"to_do_list/common/config"
	"to_do_list/module/todo/models"
	"to_do_list/module/todo/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDbConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.ToDo{})
	r := routes.SetupRouter()
	//running server
	r.Run()
}
