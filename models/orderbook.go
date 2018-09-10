package models

import "time"

// OrderType ...
type OrderType int

// OrderType types
const (
	BID OrderType = iota
	ASK
	MATCH
	CANCEL
	REFUND
	IGNORE
)

// String ...
func (o OrderType) String() string {
	switch o {
	case BID:
		return "stbid"
	case ASK:
		return "stask"
	case MATCH:
		return "match"
	case CANCEL:
		return "cancel"
	case REFUND:
		return "refund"
	case IGNORE:
		return "ignore"
	default:
		return ""
	}
}

// OrderInfo ...
type OrderInfo struct {
	Price  uint64    `json:"price"`
	Volume uint64    `json:"volume"`
	Type   OrderType `json:"type"`
}

// OrderBook ...
type OrderBook struct {
	AskRow []*OrderInfo `json:"ask"`
	BidRow []*OrderInfo `json:"bid"`
}

// UserOrderInfo ...
type UserOrderInfo struct {
	ID          int64     `json:"id"`
	OrderSymbol string    `json:"orderSymbol"`
	OrderTime   time.Time `json:"orderTime"`
	*OrderInfo
}
