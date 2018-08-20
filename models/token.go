package models

// Token ...
type Token struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	BaseSymbol      string `json:"baseSymbol"`
	Account         string `json:"account"`
	ContractAccount string `json:"contractAccount"`
	CurrentPrice    int    `json:"currentPrice"`
	PrevPrice       int    `json:"prevPrice"`
	Volume          uint   `json:"volume"`
}
