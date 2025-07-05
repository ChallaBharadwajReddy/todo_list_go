package main

import (
	controllers "gin/Controllers"
	database "gin/Database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	router := gin.Default()

	router.GET("/get_list", controllers.Get_todo)

	router.POST("/add_todo", controllers.Add_todo)

	router.GET("/", controllers.Pong)

	router.POST("/edit_todo", controllers.Edit_todo)

	router.DELETE("/delete/:id", controllers.Delete_todo)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
