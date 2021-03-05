package flexo

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"flexo/model"
)

//TODO beaucoup trop long
func (s *Server) teamReport(c *gin.Context) {
	id_str := c.Param("ID")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID isn't an int")
		return
	}

	baseMultiplier := 5

	var (
		score    int
		team     model.Team
		timeline []model.Event
		targets  []model.Target

		cat model.Category
	)

	res := s.DB.First(&team, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, "no team with that ID")
		return
	}

	res = s.DB.Where(fmt.Sprintf("%d = ANY (teams)", id)).
		Order("created_at ASC").Find(&timeline)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, "couldn't fetch timeline")
		return
	}

	targs := make(map[int64]bool)
	for _, e := range timeline {
		res := s.DB.First(&cat, e.Category)
		if res.Error != nil {
			c.JSON(http.StatusInternalServerError, "couldn't fetch data")
			return
		}
		score += cat.Multiplier

		//Get the targets
		for _, t := range e.Targets {
			targs[t] = true
		}
	}

	score *= baseMultiplier

	targsQstring := "SELECT * FROM targets WHERE"
	for i := range targs {
		targsQstring = fmt.Sprintf(" %s id = %d OR", targsQstring, i)
	}
	targsQstring += "$"
	targsQstring = strings.Replace(targsQstring, "OR$", "", -1)

	res = s.DB.Raw(targsQstring).Find(&targets)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, "couldn't fetch targets")
		return
	}

	report := model.TeamReport{
		Score:    score,
		Timeline: timeline,
		Targets:  targets,
	}

	c.JSON(http.StatusOK, report)
}
