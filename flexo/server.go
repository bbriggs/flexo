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
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/SECCDC/flexo/util"
)

type Config struct {
	DBUser string
	DBPass string
	DBAddr string
	DBName string
	DBssl  string
	Secret string
}

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func Migrate(c Config) error {
	fmt.Println("Running migrations...")
	err := util.DBinit(os.Getenv("DATABASE_URL"))
	return err
}

func Run(c Config) {
	fmt.Println("Starting Flexo...")

	var (
		db   *gorm.DB
		port string
	)

	if os.Getenv("PORT") == "" {
		port = "8080"
	} else {
		port = os.Getenv("PORT")
	}

	if os.Getenv("DATABASE_URL") != "" {
		fmt.Println("DATABASE_URL set")
		db = util.DBconnect(os.Getenv("DATABASE_URL"))
	} else {
		fmt.Printf("DATABASE_URL not set. Using other sources for configuration\n. Connecting to database %s on host %s...\n", c.DBName, c.DBAddr)
		db = util.DBconnect(util.NewConnectionString(c.DBUser, c.DBPass, c.DBAddr, c.DBName, c.DBssl))
	}

	s := Server{
		Router: gin.New(),
		DB:     db,
	}

	s.Router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthz"),
		gin.Recovery(),
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"authorization", "origin", "content-type", "accept"},
		}),
	)

	migrateErr := Migrate(c)
	if migrateErr != nil {
		fmt.Println(migrateErr)
	} else {
		fmt.Println("Migrations completed successfully!")
	}

	authorized := s.Router.Group("/")
	authorized.Use(util.SecretProvided(c.Secret))
	{
		authorized.GET("/targets", s.getTargets)
		authorized.POST("/target", s.postTarget)

		authorized.GET("/teams", s.getTeams)
		authorized.POST("/team", s.postTeam)

		authorized.GET("/categories", s.getCategories)
		authorized.POST("/category", s.postCategory)

		authorized.GET("/events", s.getEvents)
		authorized.POST("/event", s.postEvent)

		authorized.GET("/ecom", s.getEcomEvents)
		authorized.POST("/ecom", s.postEcomEvent)

		authorized.GET("/report/teams", s.allTeamsReport)
		authorized.GET("/report/team/:ID", s.teamReport)
	}

	s.Router.GET("/healthz", s.healthCheck)

	err := s.Router.Run(":" + port)
	if err != nil {
		fmt.Printf("Ran into an error: %s\n", err)
	}

	defer fmt.Println("Goodbye!")
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
