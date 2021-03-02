package flexo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"flexo/model"
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
