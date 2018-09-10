package models

import "time"

// EosdaqTx ...
type EosdaqTx struct {
	TXID          uint      `json:"-" gorm:"primary_key"`
	ID            int64     `json:"id"`
	OrderSymbol   string    `json:"orderSymbol"`
	OrderTime     time.Time `json:"orderTime"`
	TransactionID []byte    `json:"-"`

	*EOSData
}

// EOSData ...
type EOSData struct {
	// for Backend DB
	AccountName string    `json:"accountName"`
	Price       uint64    `json:"price"`
	Volume      uint64    `json:"volume"`
	Symbol      string    `json:"symbol"`
	Type        OrderType `json:"type"`
}
