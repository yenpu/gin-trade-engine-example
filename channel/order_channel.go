package channel

import (
	"github.com/yenpu/gin-trade-engine-example/domain"
	"github.com/yenpu/gin-trade-engine-example/engine"
)

var ch = make(chan domain.Order, 100)

func Send(order domain.Order) domain.Order {
	ch <- order
	return order
}

func Listen() {
	book := engine.OrderBook{
		BuyOrders:  make([]domain.Order, 0, 100),
		SellOrders: make([]domain.Order, 0, 100),
	}
	go func() {
		for {
			order := <-ch
			book.CreateOrder(order)
		}
	}()
}
