package service

import (
	"context"
	"tiffanyblue/models"
	"tiffanyblue/repository"
	"time"

	"github.com/juju/errors"
)

type symbolUsecase struct {
	obRepo     repository.OrderBookRepository
	txRepo     repository.EosdaqTxRepository
	ctxTimeout time.Duration
}

// NewSymbolService ...
func NewSymbolService(obr repository.OrderBookRepository,
	txr repository.EosdaqTxRepository,
	timeout time.Duration) SymbolService {
	return &symbolUsecase{
		obRepo:     obr,
		txRepo:     txr,
		ctxTimeout: timeout,
	}
}

// GetTickers ...
func (su symbolUsecase) GetTickers(ctx context.Context) (ts []*models.Token, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, su.ctxTimeout)
	defer cancel()

	ts, err = su.txRepo.GetTickers(innerCtx)
	if err != nil {
		return nil, errors.Annotatef(err, "GetTickers")
	}
	return ts, nil
}

// GetTicker ...
func (su symbolUsecase) GetTicker(ctx context.Context, symbol string) (ticker *models.Token, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, su.ctxTimeout)
	defer cancel()

	ticker, err = su.txRepo.GetTicker(innerCtx, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetTicker symbol[%s]", symbol)
	}
	return ticker, nil
}

// GetSymbolTxList ...
func (su symbolUsecase) GetSymbolTxList(ctx context.Context, symbol string) (txs []*models.EosdaqTx, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, su.ctxTimeout)
	defer cancel()

	txs, err = su.txRepo.GetSymbolTxList(innerCtx, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetSymbolTxList symbol[%s]", symbol)
	}
	return txs, nil
}

// GetSymbolOrderBook ...
func (su symbolUsecase) GetSymbolOrderBook(ctx context.Context, symbol string) (ob *models.OrderBook, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, su.ctxTimeout)
	defer cancel()

	obinfos, err := su.obRepo.GetOrderInfos(innerCtx, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetSymbolOrderBook symbol[%s]", symbol)
	}
	ob = &models.OrderBook{}
	for _, info := range obinfos {
		if info.Type == models.ASK {
			ob.AskRow = append(ob.AskRow, info)
		} else if info.Type == models.BID {
			ob.BidRow = append(ob.BidRow, info)
		}
	}
	return ob, nil
}
