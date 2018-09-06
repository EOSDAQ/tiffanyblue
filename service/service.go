// Package service ...
//
// This layer will act as the business process handler.
// Any process will handled here. This layer will decide, which repository layer will use.
// And have responsibility to provide data to serve into delivery.
// Process the data doing calculation or anything will done here.
//
// Service layer will accept any input from Delivery layer,
// that already sanitized, then process the input could be storing into DB ,
// or Fetching from DB ,etc.
//
// This Service layer will depends to Repository Layer
package service

import (
	"context"
	"tiffanyBlue/models"
	"tiffanyBlue/util"

	"go.uber.org/zap"
)

var mlog *zap.SugaredLogger

func init() {
	mlog, _ = util.InitLog("service", "console")
}

// ChartService ...
type ChartService interface {
	GetByID(ctx context.Context, id string) (*models.Chart, error)
}

// OrderBookService ...
type OrderBookService interface {
	GetOrderBooks(ctx context.Context, symbol string) (obs *models.OrderBook, err error)
}

// EosdaqTxService ...
type EosdaqTxService interface {
	GetTickers(ctx context.Context) (ts []*models.Token, err error)
	GetTicker(ctx context.Context, symbol string) (ticker *models.Token, err error)
	GetSymbolTxList(ctx context.Context, symbol string) (txs []*models.EosdaqTx, err error)
}

// UserService ...
type UserService interface {
	GetUserTxList(ctx context.Context, accountName string, offset int64) (txs []*models.EosdaqTx, err error)
	GetUserSymbolTxList(ctx context.Context, accountName, symbol string) (txs []*models.EosdaqTx, err error)
	GetUserOrderBook(ctx context.Context, accountName string) (obs *models.OrderBook, err error)
	GetUserSymbolOrderBook(ctx context.Context, accountName, symbol string) (obs *models.OrderBook, err error)
}
