package person

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Person struct
type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

// Repository struct
type Repository struct{}

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Person{})
}

// GetAll is get all person
func (m Repository) GetAll() ([]Person, error) {
	var p []Person
	if err := db.Find(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// CreateModel is create person model
func (m Repository) CreateModel(c *gin.Context) (Person, error) {
	var p Person
	c.BindJSON(&p)

	if err := db.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

// GetById is get a person
func (m Repository) GetById(id string) (Person, error) {
	var p Person
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

// UpdateById is update a person
func (m Repository) UpdateById(id string, c *gin.Context) (Person, error) {
	var p Person
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}

	c.BindJSON(&p)
	db.Save(&p)

	return p, nil
}

// DeleteById is delete a person
func (m Repository) DeleteById(id string) error {
	var p Person
	if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}
	return nil
}
