package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Asuforce/gogo/todo/src/model"
	"github.com/gin-gonic/gin"
)

func TaskGET(c *gin.Context) {
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

func TaskPOST(c *gin.Context) {
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

func TaskPATCH(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))

	task, err := model.TaskByID(db, uint(id))
	if err != nil {
		log.Fatal(err)
	}

	title := c.PostForm("title")
	now := time.Now()

	task.Title = title
	task.UpdatedAt = now

	err = task.Update(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(task)
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func TaskDELETE(c *gin.Context) {
	db := model.DBConnect()

	id, _ := strconv.Atoi(c.Param("id"))

	task, err := model.TaskByID(db, uint(id))
	if err != nil {
		log.Fatal(err)
	}

	err = task.Delete(db)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, "deleted")
}
