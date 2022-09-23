package main

import (
	"gin-trade-engine-example/channel"
	"gin-trade-engine-example/controller"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api", gin.BasicAuth(gin.Accounts{
		"user": "password",
	}))
	api.POST("/orders", controller.CreateOrder)
	api.GET("/buys", controller.GetBuyOrders)
	api.GET("/sells", controller.GetSellOrders)
	api.GET("/trades", controller.GetTrades)

	return r
}

func setupListener() {
	channel.Listen()
}

func main() {
	setupListener()
	r := setupRouter()
	r.Run(":8080")
}
