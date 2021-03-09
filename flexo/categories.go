package flexo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"flexo/model"
)

func (s *Server) getCategories(c *gin.Context) {
	cats, err := queryCategories(s.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Encountered an error while processing",
		})
		return
	}

	c.JSON(http.StatusOK, cats)
}

func queryCategories(db *gorm.DB) ([]model.Category, error) {
	var categories []model.Category

	res := db.Find(&categories)
	return categories, res.Error
}

func (s *Server) postCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := s.DB.Create(&category)
	if res.Error != nil {
		fmt.Println(res.Error)
		c.JSON(http.StatusInternalServerError, "Couldn't create category")
		return
	}
}
