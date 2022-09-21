package engine

import (
	"time"

	"github.com/yenpu/gin-trade-engine-example/domain"
	"github.com/yenpu/gin-trade-engine-example/util"

	"math/rand"
)

type OrderBook struct {
	BuyOrders  []domain.Order
	SellOrders []domain.Order
}

var book = OrderBook{
	BuyOrders:  make([]domain.Order, 0, 100),
	SellOrders: make([]domain.Order, 0, 100),
}

func (book *OrderBook) CreateOrder(order domain.Order) domain.Order {
	if order.OrderType == domain.Buy {
		book.buy(order)
	} else {
		book.sell(order)
	}
	return order
}

func (book *OrderBook) buy(order domain.Order) []domain.Trade {
	trades := make([]domain.Trade, 0)
	nonMatchedSellOrders := make([]domain.Order, 0)
	n := len(book.SellOrders)
	if n > 0 {
		sellOrders := book.SellOrders
		quantity := order.Quantity
		for i := 0; i < n; i++ {
			sellOrder := sellOrders[i]
			if sellOrder.Price == order.Price && sellOrder.Quantity <= quantity {
				trades = append(trades, domain.Trade{util.GetUUID(), order.ID, sellOrder.ID, order.Price, quantity})
				quantity -= sellOrder.Quantity
			} else {
				nonMatchedSellOrders = append(nonMatchedSellOrders, sellOrders[i])
			}
		}
		if quantity > 0 {
			book.BuyOrders = append(book.BuyOrders, order)
			return nil
		}
	} else {
		book.BuyOrders = append(book.BuyOrders, order)
	}

	book.SellOrders = nonMatchedSellOrders
	return trades
}

func (book *OrderBook) sell(order domain.Order) []domain.Trade {
	trades := make([]domain.Trade, 0)
	nonMatchedBuyOrders := make([]domain.Order, 0)
	n := len(book.BuyOrders)
	if n > 0 {
		buyOrders := book.BuyOrders
		quantity := order.Quantity
		for i := 0; i < n; i++ {
			buyOrder := buyOrders[i]
			if buyOrder.Price == order.Price && buyOrder.Quantity <= quantity {
				trades = append(trades, domain.Trade{util.GetUUID(), buyOrder.ID, order.ID, order.Price, quantity})
				quantity -= buyOrder.Quantity
			} else {
				nonMatchedBuyOrders = append(nonMatchedBuyOrders, buyOrders[i])
			}
		}
		if quantity > 0 {
			book.SellOrders = append(book.SellOrders, order)
			return nil
		}
	} else {
		book.SellOrders = append(book.SellOrders, order)
	}

	book.BuyOrders = nonMatchedBuyOrders
	return trades
}

func getCurrentMarketPrice() int64 {
	prices := []int64{
		100,
		110,
		123,
	}
	return prices[rand.Intn(len(prices))]
}

func FillInOrder(order domain.Order) domain.Order {
	if order.ID != "" {
		order.ID = util.GetUUID()
	}
	if order.PriceType == domain.MarketPrice {
		order.Price = getCurrentMarketPrice()
	}
	order.CreatedAt = time.Now().UnixMilli()
	return order
}
