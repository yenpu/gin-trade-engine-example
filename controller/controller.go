// Package controller is for handle the http reqeust
package controller

import (
	"net/http"

	"gin-trade-engine-example/channel"
	"gin-trade-engine-example/domain"
	"gin-trade-engine-example/engine"
	"gin-trade-engine-example/repository"

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

func GetBuyOrders(c *gin.Context) {
	repo := repository.GetBuyOrderRepo()
	c.IndentedJSON(http.StatusOK, repo.GetAll())
}

func GetSellOrders(c *gin.Context) {
	repo := repository.GetSellOrderRepo()
	c.IndentedJSON(http.StatusOK, repo.GetAll())
}

func GetTrades(c *gin.Context) {
	repo := repository.GetTradeRepo()
	c.IndentedJSON(http.StatusOK, repo.GetAll())
}
