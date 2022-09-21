package controller

import (
	"net/http"

	"github.com/yenpu/gin-trade-engine-example/channel"
	"github.com/yenpu/gin-trade-engine-example/domain"
	"github.com/yenpu/gin-trade-engine-example/engine"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order domain.Order

	if err := c.BindJSON(&order); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	order = engine.FillInOrder(order)
	go channel.Send(order)
	c.IndentedJSON(http.StatusCreated, order)
}
