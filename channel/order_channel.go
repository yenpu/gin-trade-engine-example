package channel

import (
	"gin-trade-engine-example/domain"
	"gin-trade-engine-example/engine"
)

var ch = make(chan domain.Order, 100)

func Send(order domain.Order) domain.Order {
	ch <- order
	return order
}

func Listen() {
	book := engine.InMemOrderBook{
		BuyOrders:  make([]domain.Order, 0, 100),
		SellOrders: make([]domain.Order, 0, 100),
		Trades:     make([]domain.Trade, 0, 100),
	}
	go func() {
		for {
			order := <-ch
			book.CreateOrder(order)
		}
	}()
}
