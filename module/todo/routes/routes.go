package routes

import (
	"github.com/gin-gonic/gin"
	"to_do_list/module/todo/handlers"
)

//SetupRouter untuk setup routernya bro
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", handlers.GetTodos)
		v1.GET("todo/:id", handlers.GetATodo)
		v1.POST("todo", handlers.CreateTodo)
		v1.PUT("todo/:id", handlers.EditTodo)
		v1.DELETE("todo/:id", handlers.DeleteTodo)
	}

	return r
}
