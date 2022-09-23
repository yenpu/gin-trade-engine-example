package domain

import "errors"

type OrderType string

const (
	Buy  OrderType = "buy"
	Sell           = "sell"
)

type PriceType string

const (
	MarketPrice PriceType = "market"
	LimitPrice            = "limit"
)

var orderTypeMap map[string]OrderType
var priceTypeMap map[string]PriceType

func init() {
	orderTypeMap = map[string]OrderType{
		"buy":  Buy,
		"sell": Sell,
	}
	priceTypeMap = map[string]PriceType{
		"market": MarketPrice,
		"limit":  LimitPrice,
	}
}

type Order struct {
	ID        string    `json:"id"`
	OrderType OrderType `json:"orderType" binding:"required"`
	PriceType PriceType `json:"priceType" binding:"required"`
	Price     int64     `json:"price"`
	Quantity  int64     `json:"quantity" binding:"required"`
	CreatedAt int64     `json:"createdAt"`
}

// Validate the input order
func (order *Order) Validate() (Order, error) {
	if _, exist := orderTypeMap[string(order.OrderType)]; !exist {
		return *order, errors.New(`orderType should be "buy" or "sell"`)
	}
	if _, exist := priceTypeMap[string(order.PriceType)]; !exist {
		return *order, errors.New(`priceType should be "market" or "limit"`)
	}
	return *order, nil
}

type Trade struct {
	ID       string
	BuyID    string
	SellID   string
	Price    int64
	Quantity int64
}
