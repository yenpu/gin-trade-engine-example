package engine

import (
	"gin-trade-engine-example/domain"

	"testing"
	"time"
)

func TestBuy(t *testing.T) {
	book := InMemOrderBook{
		BuyOrders:  make([]domain.Order, 0, 100),
		SellOrders: produceLimitSellOrders(),
		Trades:     make([]domain.Trade, 0, 100),
	}

	order := domain.Order{"1", "buy", "limit", 90, 1, time.Now().UnixMilli()}
	trades := book.buy(order)
	if len(trades) == 0 || trades[0].ID == "" || trades[0].BuyID != "1" || trades[0].SellID != "2" {
		t.Fatalf(`The buyId = %s, price = %d didn't match.`, order.ID, order.Price)
	}
}

func TestOneBuyMatchedTwoSells(t *testing.T) {
	book := InMemOrderBook{
		BuyOrders:  make([]domain.Order, 0, 100),
		SellOrders: produceLimitSellOrders(),
		Trades:     make([]domain.Trade, 0, 100),
	}

	order := domain.Order{"1", "buy", "limit", 98, 3, time.Now().UnixMilli()}
	trades := book.buy(order)
	if len(trades) != 2 {
		t.Fatalf(`A buy need to match with two sells.`)
	}
}

func TestSell(t *testing.T) {
	book := InMemOrderBook{
		BuyOrders:  produceLimitBuyOrders(),
		SellOrders: make([]domain.Order, 0, 100),
		Trades:     make([]domain.Trade, 0, 100),
	}

	order := domain.Order{"1", "sell", "limit", 107, 2, time.Now().UnixMilli()}
	trades := book.sell(order)
	if len(trades) == 0 || trades[0].ID == "" || trades[0].BuyID != "6" || trades[0].SellID != "1" {
		t.Fatalf(`The sellId = %s, price = %d didn't match.`, order.ID, order.Price)
	}
}

func TestOneSellMatchedThreeBuys(t *testing.T) {
	book := InMemOrderBook{
		BuyOrders:  produceLimitBuyOrders(),
		SellOrders: make([]domain.Order, 0, 100),
		Trades:     make([]domain.Trade, 0, 100),
	}

	order := domain.Order{"1", "sell", "limit", 99, 4, time.Now().UnixMilli()}
	trades := book.sell(order)
	if len(trades) != 3 {
		t.Fatalf(`A sell need to match with three buys`)
	}
}

func produceLimitSellOrders() []domain.Order {
	orders := []domain.Order{
		domain.Order{"1", "sell", "limit", 110, 3, time.Now().UnixMilli()},
		domain.Order{"2", "sell", "limit", 90, 1, time.Now().UnixMilli()},
		domain.Order{"3", "sell", "limit", 130, 1, time.Now().UnixMilli()},
		domain.Order{"4", "sell", "limit", 110, 1, time.Now().UnixMilli()},
		domain.Order{"5", "sell", "limit", 110, 1, time.Now().UnixMilli()},
		domain.Order{"6", "sell", "limit", 123, 2, time.Now().UnixMilli()},
		domain.Order{"7", "sell", "limit", 98, 1, time.Now().UnixMilli()},
		domain.Order{"8", "sell", "limit", 110, 1, time.Now().UnixMilli()},
		domain.Order{"9", "sell", "limit", 98, 2, time.Now().UnixMilli()},
		domain.Order{"10", "sell", "limit", 120, 1, time.Now().UnixMilli()},
	}
	return orders
}

func produceLimitBuyOrders() []domain.Order {
	orders := []domain.Order{
		domain.Order{"1", "buy", "limit", 111, 2, time.Now().UnixMilli()},
		domain.Order{"2", "buy", "limit", 97, 1, time.Now().UnixMilli()},
		domain.Order{"3", "buy", "limit", 120, 1, time.Now().UnixMilli()},
		domain.Order{"4", "buy", "limit", 99, 2, time.Now().UnixMilli()},
		domain.Order{"5", "buy", "limit", 113, 1, time.Now().UnixMilli()},
		domain.Order{"6", "buy", "limit", 107, 2, time.Now().UnixMilli()},
		domain.Order{"7", "buy", "limit", 100, 1, time.Now().UnixMilli()},
		domain.Order{"8", "buy", "limit", 99, 1, time.Now().UnixMilli()},
		domain.Order{"9", "buy", "limit", 112, 2, time.Now().UnixMilli()},
		domain.Order{"10", "buy", "limit", 109, 1, time.Now().UnixMilli()},
		domain.Order{"8", "buy", "limit", 99, 1, time.Now().UnixMilli()},
	}
	return orders
}
