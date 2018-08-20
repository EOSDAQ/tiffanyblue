package service

import (
	"context"
	"tiffanyBlue/models"
	"tiffanyBlue/repository"
	"time"

	"github.com/juju/errors"
)

type tickerUsecase struct {
	tickerRepo repository.TickerRepository
	ctxTimeout time.Duration
}

// NewTickerService ...
func NewTickerService(tr repository.TickerRepository,
	timeout time.Duration) TickerService {
	return &tickerUsecase{
		tickerRepo: tr,
		ctxTimeout: timeout,
	}
}

// GetTickers ...
func (tu tickerUsecase) GetTickers(ctx context.Context) (ts []*models.Ticker, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, tu.ctxTimeout)
	defer cancel()

	ts, err = tu.tickerRepo.GetTickers(innerCtx)
	if err != nil {
		return nil, errors.Annotatef(err, "GetTickers")
	}
	return ts, nil
}

// GetTicker ...
func (tu tickerUsecase) GetTicker(ctx context.Context, symbol string) (ticker *models.Ticker, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, tu.ctxTimeout)
	defer cancel()

	ticker, err = tu.tickerRepo.GetTicker(innerCtx, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetTicker symbol[%s]", symbol)
	}
	return ticker, nil
}
