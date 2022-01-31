package flexo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	pq "github.com/lib/pq"
	"gorm.io/gorm"

	"github.com/SECCDC/flexo/model"
)

// getTeams: returns a list of all the teams as a JSON array.
func (s *Server) getTeams(c *gin.Context) {
	teams, err := queryTeams(s.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Encountered an error while processing",
		})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func queryTeams(db *gorm.DB) ([]model.Team, error) {
	var teams []model.Team

	res := db.Find(&teams)
	return teams, res.Error
}

func queryTeamByID(db *gorm.DB, id int) (model.Team, error) {
	var team model.Team

	res := db.Where("team_id = ?", fmt.Sprintf("%d", id)).Find(&team)
	return team, res.Error
}

func (s *Server) postTeam(c *gin.Context) {
	var team model.Team

	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := s.DB.Create(&team)
	if res.Error != nil {
		fmt.Println(res.Error)
		c.JSON(http.StatusInternalServerError, "Couldn't create team")
		return
	}
}

func (s *Server) deleteTeam(c *gin.Context) {
	// Will there be a problem by not cleaning up events etc ?
	var b model.Team

	id_str := c.Param("ID")
	res := s.DB.Where("team_id = ?", id_str).Delete(&b)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, "Couldn't delete team")
		return
	}

	// let's update events mentionning the team we are deleting
	evs, err := queryEvents(s.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Couldn't update events regarding team deletion")
	}

	var newTeams pq.Int64Array
	for _, event := range evs {
		teams := event.Teams
		deleted_id, _ := strconv.ParseInt(id_str, 2, 64)
		for _, id := range teams {
			if id != deleted_id {
				newTeams= append(newTeams, id)
			}
		}
		s.DB.Model(&event).Update("Teams", newTeams)
	}

	c.JSON(http.StatusOK, b)
}
