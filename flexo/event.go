package flexo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/SECCDC/flexo/model"
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

	var sb strings.Builder
	
	var team model.Team
	var target model.Target
	var category model.Category

	s.DB.Where("id = ?", fmt.Sprintf("%d", event.Category)).Find(&category)

	sb.WriteString(fmt.Sprintf("New event in %s (x%d) category:\n * Teams: ", category.Name, category.Multiplier)
	for _, i := range event.Teams {
		s.DB.Where("team_id = ?", fmt.Sprintf("%d", i)).Find(&team)
		sb.WriteString(team.String())
		sb.WriteString("; ")
	}

	sb.WriteString("\n * Targets: ")
	for _, i := range event.Targets {
		s.DB.Where("id = ?", fmt.Sprintf("%d", i)).Find(&target)
		sb.WriteString(target.String())
		sb.WriteString("; ")
	}

	s.Bytebot.sendMessage(sb.String())

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

func (s *Server) computeEventValue(categoryID, baseMultiplier int) (int, error) {
	var cat model.Category
	res := s.DB.First(&cat, categoryID)
	return cat.Multiplier * baseMultiplier, res.Error
}

func (s *Server) fetchTeamTimeline(id int) ([]model.Event, error) {
	var team model.Team

	var timeline []model.Event

	res := s.DB.First(&team, id)
	if res.Error != nil {
		return nil, res.Error
	}

	res = s.DB.Where(fmt.Sprintf("%d = ANY (teams)", id)).
		Order("created_at ASC").Find(&timeline)
	if res.Error != nil {
		return nil, res.Error
	}

	return timeline, nil
}
