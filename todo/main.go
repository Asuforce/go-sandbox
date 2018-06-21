package main

import (
	"github.com/Asuforce/gogo/todo/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/tasks", controller.TaskGET)
		v1.POST("/tasks", controller.TaskPOST)
	}

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}
