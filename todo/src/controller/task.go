package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Asuforce/gogo/todo/src/model"
	"github.com/gin-gonic/gin"
)

func TasksGET(c *gin.Context) {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM task ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
	}

	tasks := []model.Task{}

	for result.Next() {
		task := model.Task{}
		var (
			id                   uint
			createdAt, updatedAt time.Time
			title                string
		)

		err = result.Scan(&id, &createdAt, &updatedAt, &title)
		if err != nil {
			log.Fatal(err)
		}

		task.ID = id
		task.CreatedAt = createdAt
		task.UpdatedAt = updatedAt
		task.Title = title
		tasks = append(tasks, task)
	}
	fmt.Println(tasks)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func TasksPOST(c *gin.Context) {
	db := model.DBConnect()

	title := c.PostForm("title")
	now := time.Now()

	task := &model.Task{
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := task.Save(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("post sent. title: %s", title)
}
