package main

import (
	"github.com/Asuforce/go-sandbox/gin-gorm/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var ctrl person.Controller

	r := gin.Default()

	p := r.Group("/people")
	{
		p.GET("", ctrl.Index)
		p.GET("/:id", ctrl.Show)
		p.POST("", ctrl.Create)
		p.PUT("/:id", ctrl.Update)
		p.DELETE("/:id", ctrl.Delete)
	}
	r.Run()
}
