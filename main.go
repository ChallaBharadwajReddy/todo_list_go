package main

import (
	controllers "gin/Controllers"
	database "gin/Database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitRedis()

	database.ConnectDatabase()

	router := gin.Default()

	router.GET("/todo", controllers.Get_todo)

	router.POST("/todo", controllers.Add_todo)

	router.GET("/", controllers.Pong)

	router.PUT("/todo", controllers.Edit_todo)

	router.DELETE("/todo/:id", controllers.Delete_todo)

	router.Run("0.0.0.0:8081")
}
