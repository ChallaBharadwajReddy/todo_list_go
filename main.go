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

	router.GET("/list", controllers.Get_todo)

	router.POST("/new-todo", controllers.Add_todo)

	router.GET("/", controllers.Pong)

	router.POST("/todo", controllers.Edit_todo)

	router.DELETE("/todo/:id", controllers.Delete_todo)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
