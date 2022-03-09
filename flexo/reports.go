package flexo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/SECCDC/flexo/model"
)

func (s *Server) generateOneTeamReport(t model.Team) (model.TeamReport, error) { // FIXME: Pass a pointer, not the whole struct here
	baseMultiplier := 20
	score := 0

	timeline, err := s.fetchTeamTimeline(t.TeamID)
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
		Team:     t,
	}

	return report, nil
}

func (s *Server) teamReport(c *gin.Context) {
	// Sanitize the input
	id_str := c.Param("ID")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID isn't an int")
		return
	}

	// Get the team struct
	team, err := queryTeamByID(s.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unable to find a team by that ID") // Is this the right return code?
		return
	}

	report, err := s.generateOneTeamReport(team)

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

	reps := []model.TeamReport{}

	for _, t := range teams {
		rep, err := s.generateOneTeamReport(t) // FIXME: Index starts at 0, teams start at 1. Papered this over in faker.
		if err != nil {
			//c.JSON(http.StatusInternalServerError, "Couldn't fetch report") // TODO: use error wrapping here
			fmt.Println("Error fetching team report")
			// We have a problem with the schema where we effectively have 2 primary keys.
			// This breaks the database when team_id != id, and team 99 isn't helping here.
			// Disabling the return until we fix the schema.
			// return
		}
		// This isn't visually confusing at all </s>
		reps = append(reps, rep) // parsec said it's easier for him if we just return an array in json
	}

	c.JSON(http.StatusOK, reps)
}
