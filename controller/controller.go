package controller

import (
	"net/http"

	"gin-trade-engine-example/channel"
	"gin-trade-engine-example/domain"
	"gin-trade-engine-example/engine"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order domain.Order

	if err := c.BindJSON(&order); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order = engine.FillInOrder(order)
	order, err := order.Validate()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go channel.Send(order)
	c.IndentedJSON(http.StatusCreated, order)
}
