package flexo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
