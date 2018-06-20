package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Asuforce/gogo/todo/src/controller"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}
