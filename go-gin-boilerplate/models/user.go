package models

import (
	"github.com/gin-gonic/gin"
	"github.com/asuforce/go-sandbox/go-gin-boilerplate/db"
)

// User model
type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

// GetAll is get all person
func (h User) GetAll() (*[]User, error) {
	db := db.GetDB()
	var user *[]User
	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// CreateModel is create person model
func (h User) CreateModel(c *gin.Context) (*User, error) {
	var user *User
	c.BindJSON(&user)

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetById is get a person
func (h User) GetById(id string) (*User, error) {
	var user *User
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

// UpdateById is update a person
func (h User) UpdateById(id string, c *gin.Context) (*User, error) {
	var user *User
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return user, err
	}

	c.BindJSON(&user)
	db.Save(&user)

	return user, nil
}

// DeleteById is delete a person
func (h User) DeleteById(id string) error {
	var user *User
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
