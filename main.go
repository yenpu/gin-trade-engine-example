package main

import (
	"github.com/yenpu/gin-trade-engine-example/channel"
	"github.com/yenpu/gin-trade-engine-example/controller"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/orders", controller.CreateOrder)

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
