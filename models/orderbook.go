package models

// OrderType ...
type OrderType int

// OrderType types
const (
	NONE OrderType = iota
	ASK
	BID
	MATCH
	CANCEL
	REFUND
	IGNORE
)

// String ...
func (o OrderType) String() string {
	switch o {
	case NONE:
		return ""
	case ASK:
		return "stask"
	case BID:
		return "stbid"
	case MATCH:
		return "match"
	case CANCEL:
		return "cancel"
	case REFUND:
		return "refund"
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
