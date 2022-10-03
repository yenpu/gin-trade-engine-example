// The package engine implement the most important business logic here, the
// engine will try to match the buy or sell depending on the order type, when
// the order mactched, the trade would be created.
package engine

import (
	"gin-trade-engine-example/domain"
	"gin-trade-engine-example/repository"
	"gin-trade-engine-example/util"

	"math/rand"
	"time"
)

type OrderBook interface {
	CreateOrder(order domain.Order) domain.Order
	GetBuyOrders() []domain.Order
	GetSellOrders() []domain.Order
	GetTrades() []domain.Trade
}

type InMemOrderBook struct {
	BuyOrders  []domain.Order
	SellOrders []domain.Order
	Trades     []domain.Trade
}

func (book *InMemOrderBook) GetBuyOrders() []domain.Order {
	return book.BuyOrders
}

func (book *InMemOrderBook) GetSellOrders() []domain.Order {
	return book.SellOrders
}

func (book *InMemOrderBook) GetTrades() []domain.Trade {
	return book.Trades
}

// CreateOrder Create the order by order type
func (book *InMemOrderBook) CreateOrder(order domain.Order) domain.Order {
	if order.OrderType == domain.Buy {
		book.buy(order)
	} else {
		book.sell(order)
	}

	return order
}

// The method will try to match the buy with sell list, it allows one by to match
// the multiple buys, and in this case the multiple trade created.
func (book *InMemOrderBook) buy(order domain.Order) []domain.Trade {
	trades := book.Trades
	nonMatchedSellOrders := make([]domain.Order, 0)
	n := len(book.SellOrders)
	if n > 0 {
		sellOrders := book.SellOrders
		quantity := order.Quantity
		for i := 0; i < n; i++ {
			sellOrder := sellOrders[i]
			if sellOrder.Price == order.Price && sellOrder.Quantity <= quantity {
				trades = append(trades, domain.Trade{util.GetUUID(), order.ID, sellOrder.ID, order.Price, sellOrder.Quantity})
				quantity -= sellOrder.Quantity
			} else {
				nonMatchedSellOrders = append(nonMatchedSellOrders, sellOrders[i])
			}
		}
		if quantity > 0 {
			book.BuyOrders = append(book.BuyOrders, order)
			book.Trades = trades
			book.updateRepos()
			return trades
		}
	} else {
		book.BuyOrders = append(book.BuyOrders, order)
	}

	book.SellOrders = nonMatchedSellOrders
	book.Trades = trades
	book.updateRepos()
	return trades
}

// The method will try to match the sell with buy list, it allows one by to match
// the multiple sells, and in this case the multiple trade created.
func (book *InMemOrderBook) sell(order domain.Order) []domain.Trade {
	trades := book.Trades
	nonMatchedBuyOrders := make([]domain.Order, 0)
	n := len(book.BuyOrders)
	if n > 0 {
		buyOrders := book.BuyOrders
		quantity := order.Quantity
		for i := 0; i < n; i++ {
			buyOrder := buyOrders[i]
			if buyOrder.Price == order.Price && buyOrder.Quantity <= quantity {
				trades = append(trades, domain.Trade{util.GetUUID(), buyOrder.ID, order.ID, order.Price, buyOrder.Quantity})
				quantity -= buyOrder.Quantity
			} else {
				nonMatchedBuyOrders = append(nonMatchedBuyOrders, buyOrders[i])
			}
		}
		if quantity > 0 {
			book.SellOrders = append(book.SellOrders, order)
			book.Trades = trades
			book.updateRepos()
			return trades
		}
	} else {
		book.SellOrders = append(book.SellOrders, order)
	}

	book.BuyOrders = nonMatchedBuyOrders
	book.Trades = trades
	book.updateRepos()
	return trades
}

// Update the data in repositories
func (book *InMemOrderBook) updateRepos() {
	buyOrderRepo := repository.GetBuyOrderRepo()
	buyOrderRepo.UpdateAll(book.BuyOrders)

	sellOrderRepo := repository.GetSellOrderRepo()
	sellOrderRepo.UpdateAll(book.SellOrders)

	tradeRepo := repository.GetTradeRepo()
	tradeRepo.UpdateAll(book.Trades)
}

// For order to get the current market price, when marketType = market
func getCurrentMarketPrice() int64 {
	prices := []int64{
		100,
		110,
		123,
	}
	return prices[rand.Intn(len(prices))]
}

func FillInOrder(order domain.Order) domain.Order {
	if order.ID == "" {
		order.ID = util.GetUUID()
	}
	if order.PriceType == domain.MarketPrice {
		order.Price = getCurrentMarketPrice()
	}
	order.CreatedAt = time.Now().UnixMilli()
	return order
}
