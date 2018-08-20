package models

// Ticker ...
type Ticker struct {
	ID              uint   `json:"id"`
	TickerName      string `json:"tickerName"`
	TokenSymbol     string `json:"tokenSymbol"`
	BaseSymbol      string `json:"baseSymbol"`
	TokenAccount    string `json:"tokenAccount"`
	ContractAccount string `json:"contractAccount"`
	CurrentPrice    int    `json:"currentPrice"`
	PrevPrice       int    `json:"prevPrice"`
	Volume          uint   `json:"volume"`
}
