package main

import (
	"github.com/Asuforce/go-sandbox/gin-gorm/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := gin.Default()
	var ctrl person.Controller
	r.GET("/people", ctrl.Index)
	r.GET("/people/:id", ctrl.Show)
	r.POST("/people", ctrl.Create)
	r.PUT("/people/:id", ctrl.Update)
	r.DELETE("/people/:id", ctrl.Delete)
	r.Run()
}
