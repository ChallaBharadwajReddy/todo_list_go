package controllers

import (
	"encoding/json"
	database "gin/Database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Get_todo(c *gin.Context) {

	cacheKey := "todos_list"

	val, err := database.RedisClient.Get(database.Ctx, cacheKey).Result()
	if err == nil {
		c.Data(http.StatusOK, "application/json", []byte(val))
		return
	}

	rows, err := database.Db.Query("SELECT id, todo FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var todos []database.Todo
	for rows.Next() {
		var todo database.Todo
		if err := rows.Scan(&todo.Id, &todo.Todo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}
		todos = append(todos, todo)
	}

	jsonData, err := json.Marshal(todos)
	if err == nil {
		database.RedisClient.Set(database.Ctx, cacheKey, jsonData, 0)
	}
	database.RedisClient.Set(database.Ctx, cacheKey, jsonData, time.Minute*5)

	c.JSON(http.StatusOK, todos)
}

func Add_todo(c *gin.Context) {
	var body database.Todo
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	_, err := database.Db.Exec("INSERT INTO todos (id, todo) VALUES ($1, $2)", body.Id, body.Todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert todo"})
		return
	}

	database.RedisClient.Del(database.Ctx, "todos_list")

	c.JSON(http.StatusOK, "Success")
}

func Edit_todo(c *gin.Context) {
	var body database.Todo
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	res, err := database.Db.Exec("UPDATE todos SET todo=$1 WHERE id=$2", body.Todo, body.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusOK, "TODO NOT FOUND")
		return
	}

	database.RedisClient.Del(database.Ctx, "todos_list")

	c.JSON(http.StatusOK, "Success")

}

func Delete_todo(c *gin.Context) {
	id := c.Param("id")

	res, err := database.Db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusOK, "TODO NOT FOUND")
		return
	}

	database.RedisClient.Del(database.Ctx, "todos_list")

	c.JSON(http.StatusOK, "Success")
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
