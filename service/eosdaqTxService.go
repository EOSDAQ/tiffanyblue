package service

import (
	"context"
	"tiffanyBlue/models"
	"tiffanyBlue/repository"
	"time"

	"github.com/juju/errors"
)

type eosdaqTxUsecase struct {
	txRepo     repository.EosdaqTxRepository
	ctxTimeout time.Duration
}

// NewEosdaqTxService ...
func NewEosdaqTxService(txr repository.EosdaqTxRepository,
	timeout time.Duration) EosdaqTxService {
	return &eosdaqTxUsecase{
		txRepo:     txr,
		ctxTimeout: timeout,
	}
}

// GetTickers ...
func (eu eosdaqTxUsecase) GetTickers(ctx context.Context) (ts []*models.Token, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, eu.ctxTimeout)
	defer cancel()

	ts, err = eu.txRepo.GetTickers(innerCtx)
	if err != nil {
		return nil, errors.Annotatef(err, "GetTickers")
	}
	return ts, nil
}

// GetTicker ...
func (eu eosdaqTxUsecase) GetTicker(ctx context.Context, symbol string) (ticker *models.Token, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, eu.ctxTimeout)
	defer cancel()

	ticker, err = eu.txRepo.GetTicker(innerCtx, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetTicker symbol[%s]", symbol)
	}
	return ticker, nil
}

func (eu eosdaqTxUsecase) GetSymbolTxList(ctx context.Context, symbol string) (txs []*models.EosdaqTx, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, eu.ctxTimeout)
	defer cancel()

	txs, err = eu.txRepo.GetSymbolTxList(innerCtx, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetSymbolTxList symbol[%s]", symbol)
	}
	return txs, nil
}
