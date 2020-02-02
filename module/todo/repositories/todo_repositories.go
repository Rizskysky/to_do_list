package repositories

import (
	"fmt"
	"to_do_list/db"
	"to_do_list/module/todo/models"

	_ "github.com/go-sql-driver/mysql" //mysql driver
)

//GetAllTodo dapetin semua todonya bro
func GetAllTodo() ([]models.ToDo, error) {
	var todoList []models.ToDo
	if err := db.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}

//CreateATodo is function for buat bro
func CreateATodo(todo models.ToDo) (models.ToDo, error) {
	if err := db.DB.Create(&todo).Error; err != nil {
		return models.ToDo{}, err
	}
	return todo, nil
}

//GetATodo is function for dapetin 1 bro
func GetATodo(id string) (models.ToDo, error) {
	var aTodo models.ToDo
	if err := db.DB.Where("id = ?", id).First(&aTodo).Error; err != nil {
		fmt.Println(err)
		return models.ToDo{}, err
	}
	return aTodo, nil
}

//UpdateTodo is function for update bro
func UpdateTodo(todo models.ToDo, id string) (models.ToDo, error) {
	db.DB.Save(&todo)
	return todo, nil
}

//DeleteTodo is function for delete bro
func DeleteTodo(todo models.ToDo, id string) (models.ToDo, error) {
	if err := db.DB.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return models.ToDo{}, err
	}
	return todo, nil
}
