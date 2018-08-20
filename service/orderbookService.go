package service

import (
	"context"
	"tiffanyBlue/models"
	"tiffanyBlue/repository"
	"time"

	"github.com/juju/errors"
)

type orderBookUsecase struct {
	orderBookRepo repository.OrderBookRepository
	symbolMap     map[string]string
	ctxTimeout    time.Duration
}

// NewOrderBookService ...
func NewOrderBookService(obr repository.OrderBookRepository,
	tickerRepo repository.TickerRepository,
	timeout time.Duration) OrderBookService {

	tickers, err := tickerRepo.GetTickers(context.Background())
	if err != nil {
		mlog.Errorw("NewOrderBookService", "err", err)
		return nil
	}
	obu := &orderBookUsecase{
		orderBookRepo: obr,
		ctxTimeout:    timeout,
		symbolMap:     make(map[string]string),
	}
	for _, t := range tickers {
		obu.symbolMap[t.TokenSymbol] = t.ContractAccount
	}
	return obu
}

// GetOrderBooks ...
func (obu orderBookUsecase) GetOrderBooks(ctx context.Context, symbol string) (ob *models.OrderBook, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, obu.ctxTimeout)
	defer cancel()

	var obinfos []*models.OrderInfo
	if contract, ok := obu.symbolMap[symbol]; !ok {
		mlog.Errorw("GetOrderBooks Invalid symbol", "symbol", symbol)
		return nil, errors.Errorf("GetOrderBooks Symbol[%s]", symbol)
	} else {
		obinfos, err = obu.orderBookRepo.GetOrderInfos(innerCtx, contract)
	}
	if err != nil {
		return nil, errors.Annotatef(err, "GetOrderBooks Symbol[%s]", symbol)
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
