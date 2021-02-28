package flexo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"flexo/model"
)

func (s *Server) getTargets(c *gin.Context) {
	targets, err := queryTargets(s.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Encountered an error while processing",
		})
		return
	}

	c.JSON(http.StatusOK, targets)
}

func queryTargets(db *gorm.DB) ([]model.Target, error) {
	var targets []model.Target

	res := db.Find(&targets)
	return targets, res.Error
}
