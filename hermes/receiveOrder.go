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
	"time"

	"hermes/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createOrder(db *gorm.DB, productIDs []int) (model.Order, error) {
	var products []model.Product

	// Lookup the product and make sure it exists
	q1 := db.Find(&products, productIDs)
	if q1.Error != nil {
		fmt.Println(q1.Error)
		return model.Order{}, q1.Error
	}

	//TODO: handle requested products that are not in the DB
	// Create and insert the order
	order := model.Order{
		TimeReceived: time.Now(),
	}
	q2 := db.Create(&order)
	if q2.Error != nil {
		fmt.Println(q2.Error)
		return order, q2.Error
	}

	for _, product := range products {
		res := db.Create(&model.OrderMapping{
			OrderID:   order.ID,
			ProductID: product.ID,
		})
		if res.Error != nil {
			fmt.Println(res.Error)
		}
	}
	return order, q2.Error
}

func (s *Server) receiveOrder(c *gin.Context) {
	var request model.OrderRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "Order request is malformed. Please try again.",
		})
		return
	}

	order, err := createOrder(s.DB, request.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"reason": "Failed to create order.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "created",
		"id":     order.ID,
	})
}
