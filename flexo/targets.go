package flexo

import (
	"fmt"
	"net/http"
	"strings"

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

func (s *Server) getTargetList(ids []int) ([]model.Target, error) {
	var targets []model.Target

	targsQstring := "SELECT * FROM targets WHERE"
	for i := range ids {
		targsQstring = fmt.Sprintf(" %s id = %d OR", targsQstring, i)
	}
	targsQstring += "$"
	targsQstring = strings.Replace(targsQstring, "OR$", "", -1)

	res := s.DB.Raw(targsQstring).Find(&targets)

	return targets, res.Error
}

func (s *Server) postTarget(c *gin.Context) {
	var target model.Target

	if err := c.ShouldBindJSON(&target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := s.DB.Create(&target)
	if res.Error != nil {
		fmt.Println(res.Error)
		c.JSON(http.StatusInternalServerError, "Couldn't create target")
		return
	}
}
