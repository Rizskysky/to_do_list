package routes

import (
	"net/http"
	"to_do_list/module/todo/handlers"

	"github.com/gin-gonic/gin"
)

//SetupRouter untuk setup routernya bro
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//404 not found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "data": nil, "message": "Page not found"})
	})

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
