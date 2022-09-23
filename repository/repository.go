// Package repository implement the storage of domain objects, in this example
// project, use the in-memory storage for simplicity.
package repository

import (
	"gin-trade-engine-example/domain"
)

var buyOrderRepo BuyOrderRepository = BuyOrderRepository{}
var sellOrderRepo SellOrderRepository = SellOrderRepository{}
var tradeRepo TradeRepository = TradeRepository{}

var buyOrders = make([]domain.Order, 0, 100)
var sellOrders = make([]domain.Order, 0, 100)
var repoTrades = make([]domain.Trade, 0, 100)

type BuyOrderRepository struct {
}

type SellOrderRepository struct {
}

type TradeRepository struct {
}

func (repo *BuyOrderRepository) GetAll() []domain.Order {
	return buyOrders
}

func (repo *BuyOrderRepository) UpdateAll(orders []domain.Order) []domain.Order {
	buyOrders = orders
	return buyOrders
}

func (repo *SellOrderRepository) GetAll() []domain.Order {
	return sellOrders
}

func (repo *SellOrderRepository) UpdateAll(orders []domain.Order) []domain.Order {
	sellOrders = orders
	return sellOrders
}

func (repo *TradeRepository) GetAll() []domain.Trade {
	return repoTrades
}

func (repo *TradeRepository) UpdateAll(trades []domain.Trade) []domain.Trade {
	repoTrades = trades
	return repoTrades
}

func GetBuyOrderRepo() BuyOrderRepository {
	return buyOrderRepo
}

func GetSellOrderRepo() SellOrderRepository {
	return sellOrderRepo
}

func GetTradeRepo() TradeRepository {
	return tradeRepo
}
