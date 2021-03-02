package flexo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"flexo/model"
)

func (s *Server) postEvent(c *gin.Context) {
	var event model.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := s.DB.Create(&event)
	if res.Error != nil {
		fmt.Println(res.Error)
		c.JSON(http.StatusInternalServerError, "Couldn't create event")
		return
	}
}

func (s *Server) getEvents(c *gin.Context) {
	events, err := queryEvents(s.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Encountered an error while processing",
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

func queryEvents(db *gorm.DB) ([]model.Event, error) {
	var events []model.Event

	res := db.Find(&events)
	return events, res.Error
}
