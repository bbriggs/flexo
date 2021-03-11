package flexo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/SECCDC/flexo/model"
)

func (s *Server) teamReport(c *gin.Context) {
	id_str := c.Param("ID")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID isn't an int")
		return
	}

	baseMultiplier := 20
	score := 0

	timeline, err := s.fetchTeamTimeline(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "couldn't fetch timeline")
		return
	}

	targs := []int{}
	for _, e := range timeline {
		sc, err := s.computeEventValue(e.Category, baseMultiplier)
		if err != nil {
			fmt.Printf("Couldn't compute event %d's value\n", e.ID)
		}
		score += sc

		for _, t := range e.Targets {
			targs = append(targs, int(t))
		}
	}

	targets, err := s.getTargetList(targs)

	report := model.TeamReport{
		Score:    score,
		Timeline: timeline,
		Targets:  targets,
	}

	c.JSON(http.StatusOK, report)
}
