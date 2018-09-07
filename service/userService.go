package service

import (
	"context"
	"tiffanyBlue/models"
	"tiffanyBlue/repository"
	"time"

	"github.com/juju/errors"
)

type userUsecase struct {
	obRepo     repository.OrderBookRepository
	txRepo     repository.EosdaqTxRepository
	ctxTimeout time.Duration
}

// NewUserService ...
func NewUserService(obr repository.OrderBookRepository,
	txr repository.EosdaqTxRepository,
	timeout time.Duration) UserService {
	return &userUsecase{
		obRepo:     obr,
		txRepo:     txr,
		ctxTimeout: timeout,
	}
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

// GetUserSymbolOrderInfos ...
func (uc *userUsecase) GetUserSymbolOrderInfos(ctx context.Context, accountName, symbol string) (obs []*models.OrderInfo, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	obs, err = uc.obRepo.GetUserSymbolOrderInfos(innerCtx, accountName, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetUserSymbolOrderBook account[%s] symbol[%s]", accountName, symbol)
	}
	return obs, nil
}

// GetUserTxList ...
func (uc *userUsecase) GetUserTxList(ctx context.Context, accountName string, page uint) (txs []*models.EosdaqTx, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	txs, err = uc.txRepo.GetUserTxList(innerCtx, accountName, page)
	if err != nil {
		return nil, errors.Annotatef(err, "GetUserTxList account[%s]", accountName)
	}
	return txs, nil
}

// GetUserOrderInfos ...
func (uc *userUsecase) GetUserOrderInfos(ctx context.Context, accountName string) (obs []*models.OrderInfo, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	obs, err = uc.obRepo.GetUserOrderInfos(innerCtx, accountName)
	if err != nil {
		return nil, errors.Annotatef(err, "GetUserOrderBook account[%s]", accountName)
	}

	return obs, nil
}
