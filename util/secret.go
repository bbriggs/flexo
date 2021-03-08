package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// A stupid simple middleware that checks that a request has a
// Authorization: Basic $SECRET
// header
func SecretProvided(secret string) gin.HandlerFunc {
	expectedValue := fmt.Sprintf("Bearer %s", secret)
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != expectedValue {
			fmt.Printf("Wrong auth token provided: %s\n", c.GetHeader("Authorization"))
			c.JSON(http.StatusUnauthorized, "Wrong secret")
			c.Abort()
		}

		c.Next()
	}
}
