package controllers

import (
	"encoding/json"
	"fmt"
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

	todos, err := database.Get_list()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch todos"})
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

	body.Id = fmt.Sprintf("%d", time.Now().UnixNano())

	err := database.Insert_todo(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to insert todo"})
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

	rowsAffected, err := database.Edit_todo(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusOK, "TODO NOT FOUND")
		return
	}

	database.RedisClient.Del(database.Ctx, "todos_list")

	c.JSON(http.StatusOK, "Success")

}

func Delete_todo(c *gin.Context) {
	id := c.Param("id")

	rowsAffected, err := database.Delete_todo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}

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
