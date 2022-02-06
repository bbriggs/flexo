package flexo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/SECCDC/flexo/model"
)

func (s *Server) postEcomEvent(c *gin.Context) {
	var event model.EcomEvent

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

	s.Bytebot.sendMessage(fmt.Sprintf("New ecom event: %s", event))
}

func (s *Server) getEcomEvents(c *gin.Context) {
	events, err := queryEcomEvents(s.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Encountered an error while processing",
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

func queryEcomEvents(db *gorm.DB) ([]model.EcomEvent, error) {
	var events []model.EcomEvent

	res := db.Find(&events)
	return events, res.Error
}
