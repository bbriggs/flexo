package flexo

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"flexo/model"
)

// event : Takes an event's category ID, a description string,
// a comma separated list of ints for the involved team's IDs,
// and another for the target's.
// For instance, one could create an event using httpie with this command:
// `http --form POST flexo_1/event category=1 description="a sample event" teams=1,2,3 targets=3,2,1`
func (s *Server) event(c *gin.Context) {
	// The values aren't passed correctly for some reason.
	cat, err := s.extractCategory(c.PostForm("category"))
	if err != nil {
		fmt.Println(err)
		//TODO More precise?
		c.JSON(http.StatusBadRequest, "Malformed category")
		return
	}

	teams, err := s.extractTeams(c.PostForm("teams"))
	if err != nil {
		fmt.Println(err)
		//TODO More precise?
		c.JSON(http.StatusBadRequest, "Malformed teams field")
		return
	}

	targets, err := s.extractTargets(c.PostForm("targets"))
	if err != nil {
		fmt.Println(err)
		//TODO More precise?
		c.JSON(http.StatusBadRequest, "Malformed targets field")
		return
	}

	event := model.Event{
		Targets:     targets,
		Teams:       teams,
		Category:    cat,
		Description: c.PostForm("description"), // TODO SQL injections ???
	}

	res := s.DB.Create(&event)
	if res.Error != nil {
		fmt.Println(res.Error)
		c.JSON(http.StatusInternalServerError, "Couldn't create event")
	}
}

func (s *Server) extractTeams(value string) ([]model.Team, error) {
	ids, err := extractListOfIDS(value)
	if err != nil {
		return nil, err
	}

	var ret []model.Team

	res := s.DB.Find(&ret, ids)

	return ret, res.Error
}

func (s *Server) extractTargets(value string) ([]model.Target, error) {
	ids, err := extractListOfIDS(value)
	if err != nil {
		return nil, err
	}

	var ret []model.Target

	res := s.DB.Find(&ret, ids)

	return ret, res.Error
}

func extractListOfIDS(value string) ([]int, error) {
	ids := []int{}

	for _, c := range strings.Split(",", value) {
		id, err := strconv.Atoi(c)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

// extractCategory extracts the category and returns it.
func (s *Server) extractCategory(value string) (model.Category, error) {
	var cat model.Category

	// supposed to escape for sql injections automatically?
	res := s.DB.First(&cat, value)

	return cat, res.Error
}
