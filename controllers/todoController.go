package controllers

import (
	"github/loa212/go-todo/initializers"
	"github/loa212/go-todo/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	// result := initializers.DB.Find(&todos)

	//get logged user id
	user, _ := c.Get("user")

	userI, _ := user.(models.User)

	// get only logged in user's todos
	result := initializers.DB.Where("user_id = ?", userI.ID ).Find(&todos)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error getting posts",
		})
		return
	}

	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func CreateTodo(c *gin.Context) {
	//get data of req.body
	var body models.Todo

	c.BindJSON(&body)

	//get user id
	user, _ := c.Get("user")
	Iuser, _ := user.(models.User)
	
	//create post
	todo := models.Todo{Title: body.Title, Done: false, UserID: Iuser.ID}

	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error creating todo",
		})
		return
	}

	//return post
	c.JSON(200, gin.H{
		"message": "todo created",
		"todo": todo,
	})
	
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var body models.Todo
	c.BindJSON(&body)

	var todo models.Todo
	result := initializers.DB.First(&todo, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error getting todo with id=" + id,
		})
		return
	}

	todo.Title = body.Title
	todo.Done = body.Done

	result = initializers.DB.Save(&todo)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error updating todo with id=" + id,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "todo updated",
		"todo": todo,
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	result := initializers.DB.First(&todo, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error getting todo with id=" + id,
		})
		return
	}

	result = initializers.DB.Delete(&todo)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error deleting todo with id=" + id,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "todo deleted",
	})
}