package service

import (
	"context"
	"tiffanyBlue/models"
	"tiffanyBlue/repository"
	"time"

	"github.com/juju/errors"
)

type tickerUsecase struct {
	tokenRepo  repository.TokenRepository
	ctxTimeout time.Duration
}

// NewTickerService ...
func NewTickerService(tr repository.TokenRepository,
	timeout time.Duration) TickerService {
	return &tickerUsecase{
		tokenRepo:  tr,
		ctxTimeout: timeout,
	}
}

// GetTickers ...
func (tu tickerUsecase) GetTickers(ctx context.Context) (ts []*models.Token, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, tu.ctxTimeout)
	defer cancel()

	ts, err = tu.tokenRepo.GetTokens(innerCtx)
	if err != nil {
		return nil, errors.Annotatef(err, "GetTickers")
	}
	return ts, nil
}

// GetTicker ...
func (tu tickerUsecase) GetTicker(ctx context.Context, symbol string) (ticker *models.Token, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, tu.ctxTimeout)
	defer cancel()

	ticker, err = tu.tokenRepo.GetToken(innerCtx, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetTicker symbol[%s]", symbol)
	}
	return ticker, nil
}
