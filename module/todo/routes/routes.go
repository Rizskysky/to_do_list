package routes

import (
	"to_do_list/module/todo/handlers"
	"to_do_list/module/todo/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var todoRepo repositories.Todo

//SetupRouter untuk setup routernya bro
func SetupRouter(db *gorm.DB) *gin.Engine {
	todoRepo = repositories.NewTodoRepositories(db)
	TodoHandler := handlers.TodoHandler(todoRepo)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", TodoHandler.GetTodos)
		v1.GET("todo/:id", TodoHandler.GetATodo)
		v1.POST("todo", TodoHandler.CreateTodo)
		v1.PUT("todo/:id", TodoHandler.EditTodo)
		v1.DELETE("todo/:id", TodoHandler.DeleteTodo)
	}

	return r
}
