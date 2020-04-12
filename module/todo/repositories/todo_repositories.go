package repositories

import (
	"fmt"
	"to_do_list/module/todo/models"

	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/jinzhu/gorm"
)

//NewTodoRepositories init function
func NewTodoRepositories(DB *gorm.DB) Todo {
	return &TodoRepository{
		db: DB,
	}

}

//GetAllTodo dapetin semua todonya bro
func (repo *TodoRepository) GetAllTodo() ([]models.ToDo, error) {
	var todoList []models.ToDo
	if err := repo.db.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}

//CreateATodo is function for buat bro
func (repo *TodoRepository) CreateATodo(todo models.ToDo) (models.ToDo, error) {
	if err := repo.db.Create(&todo).Error; err != nil {
		return models.ToDo{}, err
	}
	return todo, nil
}

//GetATodo is function for dapetin 1 bro
func (repo *TodoRepository) GetATodo(id string) (models.ToDo, error) {
	var aTodo models.ToDo
	if err := repo.db.Where("id = ?", id).First(&aTodo).Error; err != nil {
		fmt.Println(err)
		return models.ToDo{}, err
	}
	return aTodo, nil
}

//UpdateTodo is function for update bro
func (repo *TodoRepository) UpdateTodo(todo models.ToDo, id string) (models.ToDo, error) {
	repo.db.Save(&todo)
	return todo, nil
}

//DeleteTodo is function for delete bro
func (repo *TodoRepository) DeleteTodo(todo models.ToDo, id string) (models.ToDo, error) {
	if err := repo.db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return models.ToDo{}, err
	}
	return todo, nil
}
