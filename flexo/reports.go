package flexo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/SECCDC/flexo/model"
)

func (s *Server) generateOneTeamReport(id int) (model.TeamReport, error){
	baseMultiplier := 20
	score := 0

	timeline, err := s.fetchTeamTimeline(id)
	if err != nil {
		return model.TeamReport{}, err
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

	return report, nil
}

func (s *Server) teamReport(c *gin.Context) {
	id_str := c.Param("ID")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID isn't an int")
		return
	}

	report, err := s.generateOneTeamReport(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "couldn't fetch timeline")
		return
	}

	c.JSON(http.StatusOK, report)
}

func (s *Server) allTeamsReport(c *gin.Context) {
	// The output is not sorted. Is it much of a problem ?
	teams, err := queryTeams(s.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Couldn't build report")
		return
	}

	reps := make([]model.TeamReport, 0)

	for _, t := range teams {
		rep, err := s.generateOneTeamReport(int(t.ID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Couldn't fetch report")
			return
		}
		reps = append(reps, rep)
	}

	c.JSON(http.StatusOK, reps)
}
