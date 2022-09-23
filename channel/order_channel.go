// Package channel implement the producer and receiver of message queue
package channel

import (
	"fmt"
	"gin-trade-engine-example/domain"
	"gin-trade-engine-example/engine"
	"gin-trade-engine-example/repository"
	"time"
)

// Use the bufferred channel for message queue
var ch = make(chan domain.Order, 100)

func Send(order domain.Order) domain.Order {
	ch <- order
	return order
}

func Listen() {
	buyRepo := repository.GetBuyOrderRepo()
	buyOrders := buyRepo.GetAll()

	sellRepo := repository.GetSellOrderRepo()
	sellOrders := sellRepo.GetAll()

	tradeRepo := repository.GetTradeRepo()
	trades := tradeRepo.GetAll()

	book := engine.InMemOrderBook{
		BuyOrders:  buyOrders,
		SellOrders: sellOrders,
		Trades:     trades,
	}

	go func() {
		for {
			order := <-ch
			fmt.Println("Receive message")
			book.CreateOrder(order)
			time.Sleep(8 * time.Second)
		}
	}()
}
