package person

import (
	"fmt"

	"github.com/Asuforce/go-sandbox/gin-gorm/models"
	"github.com/gin-gonic/gin"
)

// Controller struct
type Controller struct{}

// Index action
func (pc Controller) Index(c *gin.Context) {
	var repo person.Repository
	p, err := repo.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, p)
}

// Create action
func (pc Controller) Create(c *gin.Context) {
	var repo person.Repository
	p, err := repo.CreateModel(c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}
	c.JSON(201, p)
}

// Show action
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var repo person.Repository
	p, err := repo.GetById(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, p)
}

// Update action
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var repo person.Repository
	p, err := repo.UpdateById(id, c)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}
	c.JSON(200, p)
}

// Delete action
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var repo person.Repository

	if err := repo.DeleteById(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	}
	c.JSON(204, gin.H{"id #" + id: "deleted"})
}
