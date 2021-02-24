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

package hermes

import (
	"fmt"
	"net/http"

	"hermes/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getProductsFromDB(db *gorm.DB) ([]model.Product, error) {
	var products []model.Product

	res := db.Find(&products)
	if res.Error != nil {
		fmt.Println(res.Error)
	}

	return products, res.Error
}

func (s *Server) listProducts(c *gin.Context) {
	products, err := getProductsFromDB(s.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "Unable to retrieve product list. Please try again later.",
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
