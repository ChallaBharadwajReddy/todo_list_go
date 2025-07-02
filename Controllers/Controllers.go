package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id   string `json:"id"`
	Todo string `json:"todo"`
}

var todo_list = []Todo{}

func Get_todo(c *gin.Context) {
	c.JSON(http.StatusOK, todo_list)
}

func Add_todo(c *gin.Context) {
	var body Todo

	c.Bind(&body)

	todo_list = append(todo_list, body)

	c.JSON(http.StatusOK, "Success")
}

func Edit_todo(c *gin.Context) {
	var body Todo

	c.Bind(&body)

	for i := range todo_list {
		if todo_list[i].Id == body.Id {
			todo_list[i].Todo = body.Todo
			c.JSON(http.StatusOK, "Success")
			return
		}
	}
	c.JSON(http.StatusOK, "TODO NOT FOUND")
}

func Delete_todo(c *gin.Context) {
	var id string = c.Param("id")

	for i := range todo_list {
		if todo_list[i].Id == id {
			todo_list = append(todo_list[:i], todo_list[i+1:]...)
			c.JSON(http.StatusOK, "Success")
			return
		}
	}
	c.JSON(http.StatusOK, "TODO NOT FOUND")
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
