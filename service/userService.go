package service

import (
	"context"
	"tiffanyBlue/models"
	"tiffanyBlue/repository"
	"time"

	"github.com/juju/errors"
)

type userUsecase struct {
	txRepo     repository.EosdaqTxRepository
	obRepo     repository.OrderBookRepository
	ctxTimeout time.Duration
}

// NewUserService ...
func NewUserService(txr repository.EosdaqTxRepository,
	obr repository.OrderBookRepository,
	timeout time.Duration) UserService {
	return &userUsecase{
		txRepo:     txr,
		obRepo:     obr,
		ctxTimeout: timeout,
	}
}

// GetUserTxList ...
func (uc *userUsecase) GetUserTxList(ctx context.Context, accountName string, offset int64) (txs []*models.EosdaqTx, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	txs, err = uc.txRepo.GetUserTxList(innerCtx, accountName, offset)
	if err != nil {
		return nil, errors.Annotatef(err, "GetUserTxList account[%s]", accountName)
	}
	return txs, nil
}

// GetUserSymbolTxList ...
func (uc *userUsecase) GetUserSymbolTxList(ctx context.Context, accountName, symbol string) (txs []*models.EosdaqTx, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	txs, err = uc.txRepo.GetUserSymbolTxList(innerCtx, accountName, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetUserSymbolTxList account[%s] symbol[%s]", accountName, symbol)
	}
	return txs, nil
}

// GetUserOrderBook ...
func (uc *userUsecase) GetUserOrderBook(ctx context.Context, accountName string) (obs *models.OrderBook, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	obinfos, err := uc.obRepo.GetUserOrderInfos(innerCtx, accountName)
	if err != nil {
		return nil, errors.Annotatef(err, "GetUserOrderBook account[%s]", accountName)
	}

	obs = ConvertOrderBook(obinfos)
	return obs, nil
}

// GetUserSymbolOrderBook ...
func (uc *userUsecase) GetUserSymbolOrderBook(ctx context.Context, accountName, symbol string) (obs *models.OrderBook, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	obinfos, err := uc.obRepo.GetUserSymbolOrderInfos(innerCtx, accountName, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetUserSymbolOrderBook account[%s] symbol[%s]", accountName, symbol)
	}
	obs = ConvertOrderBook(obinfos)
	return obs, nil
}
