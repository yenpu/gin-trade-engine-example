package channel

import (
	"gin-trade-engine-example/domain"
	"gin-trade-engine-example/engine"
	"gin-trade-engine-example/repository"
)

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
			book.CreateOrder(order)
		}
	}()
}
