package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"to_do_list/common/config"
)

//GetAllTodo dapetin semua todonya bro
func GetAllTodo(todo *[]ToDo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

//CreateATodo is function for buat bro
func CreateATodo(todo *ToDo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

//GetATodo is function for dapetin 1 bro
func GetATodo(todo *ToDo, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

//UpdateTodo is function for update bro
func UpdateTodo(todo *ToDo, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

//DeleteTodo is function for delete bro
func DeleteTodo(todo *ToDo, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
