package main

import (
	"github/loa212/go-todo/controllers"
	"github/loa212/go-todo/initializers"
	"github/loa212/go-todo/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDatabase()

	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Signin)
	//Auth required routes
	r.GET("/validate", middlewares.RequireAuth, controllers.Validate)
	r.DELETE("/user/:id", middlewares.RequireAuth, controllers.DeleteUser)
	//Todo routes
	r.GET("/todos", middlewares.RequireAuth, controllers.GetTodos)
	r.POST("/createTodo", middlewares.RequireAuth, controllers.CreateTodo)
	r.PUT("/updateTodo/:id", middlewares.RequireAuth, controllers.UpdateTodo)
	r.DELETE("/deleteTodo/:id", middlewares.RequireAuth, controllers.DeleteTodo)

	//run server
	r.Run() // listen and serve on 0.0.0.0:3000
}