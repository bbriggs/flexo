/*
Copyright © 2021 Bren 'fraq' Briggs (code@fraq.io)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package flexo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"flexo/util"
)

type Config struct {
	DBUser string
	DBPass string
	DBAddr string
	DBName string
}

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func Migrate(c Config) {
	fmt.Println("Running migrations...")
	err := util.DBinit(c.DBUser, c.DBPass, c.DBAddr, c.DBName, false) //TODO SSL connection to the DB option
	if err != nil {
		fmt.Println("Encountered errors while migrating:")
		fmt.Println(err)
		return
	}

	fmt.Println("Migrations executed successfully!")
}

func Run(c Config) {
	fmt.Println("Starting Flexo...")
	s := Server{
		Router: gin.Default(),
		DB:     util.DBconnect(c.DBUser, c.DBPass, c.DBAddr, c.DBName, false), //TODO SSL connection to the DB option
	}

	s.Router.GET("/healthz", s.healthCheck)
	s.Router.GET("/targets", s.getTargets)
	s.Router.GET("/teams", s.getTeams)
	s.Router.GET("/categories", s.getCategories)
	s.Router.POST("/event", s.event)
	s.Router.Run()

	defer fmt.Println("Goodbye!")
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
