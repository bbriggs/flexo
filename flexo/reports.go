package flexo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"flexo/model"
)

func (s *Server) teamReport(c *gin.Context) {
	id := c.Param("ID")

	var team model.Team

	res := s.DB.First(&team, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, "no team with that ID")
		return
	}

	score, err := s.getScore(team)
	if err != nil {
		//TODO
	}

	report := model.TeamReport{
		Score: score,
	}

	c.JSON(http.StatusOK, report)
}

func (s *Server) getScore(t model.Team) (int, error) {
	var events, allEvents []model.Event

	// TODO There is probably something to be done with gorm.
	// sqlQuery := fmt.Sprintf("%d = ANY (teams)", t.ID)

	res := s.DB.Find(&allEvents)
	if res.Error != nil {
		fmt.Println("Couldn't query events")
		return -1, res.Error
	}

	for _, e := range allEvents {
		for _, teamID := range e.Teams {
			if uint(teamID) == t.ID {
				events = append(events, e)
			}
		}

	}

	return s.scoreFromEvents(events)
}

func (s *Server) scoreFromEvents(ev []model.Event) (int, error) {
	score := 0
	//TODO make this configurable/more visible
	constantMultiplier := 5

	for _, e := range ev {
		s, err := s.getCategoryMultiplier(e)
		if err != nil {
			fmt.Println("Couldn't get event multiplier")
			return -1, err

		} else {
			score += s * constantMultiplier
		}
	}

	return score, nil
}

func (s *Server) getCategoryMultiplier(event model.Event) (int, error) {
	var c model.Category
	res := s.DB.First(&c, event.Category)

	return c.Multiplier, res.Error
}
