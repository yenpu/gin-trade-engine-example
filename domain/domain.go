package domain

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

type Order struct {
	ID        string    `json:"id"`
	OrderType OrderType `json:"orderType" binding:"required"`
	PriceType PriceType `json:"priceType" binding:"required"`
	Price     int64     `json:"price"`
	Quantity  int64     `json:"quantity" binding:"required"`
	CreatedAt int64     `json:"createdAt"`
}

type Trade struct {
	ID       string
	BuyID    string
	SellID   string
	Price    int64
	Quantity int64
}
