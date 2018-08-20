package models

// OrderType ...
type OrderType int

// OrderType types
const (
	ASK OrderType = iota
	BID
)

// String ...
func (o OrderType) String() string {
	switch o {
	case ASK:
		return "stask"
	case BID:
		return "stbid"
	default:
		return "tx"
	}
}

// OrderInfo ...
type OrderInfo struct {
	Price  int       `json:"price"`
	Volume int       `json:"volume"`
	Type   OrderType `json:"type"`
}

// OrderBook ...
type OrderBook struct {
	AskRow []*OrderInfo `json:"ask"`
	BidRow []*OrderInfo `json:"bid"`
}
