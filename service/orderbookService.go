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
	ctxTimeout    time.Duration
}

// NewOrderBookService ...
func NewOrderBookService(obr repository.OrderBookRepository,
	timeout time.Duration) OrderBookService {

	return &orderBookUsecase{
		orderBookRepo: obr,
		ctxTimeout:    timeout,
	}
}

// GetOrderBooks ...
func (obu orderBookUsecase) GetOrderBooks(ctx context.Context, symbol string) (ob *models.OrderBook, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, obu.ctxTimeout)
	defer cancel()

	obinfos, err := obu.orderBookRepo.GetOrderInfos(innerCtx, symbol)
	if err != nil {
		return nil, errors.Annotatef(err, "GetOrderBooks Symbol[%s]", symbol)
	}

	ob = ConvertOrderBook(obinfos)
	return ob, nil
}

func ConvertOrderBook(obinfos []*models.OrderInfo) (ob *models.OrderBook) {
	ob = &models.OrderBook{}
	for _, info := range obinfos {
		if info.Type == models.ASK {
			ob.AskRow = append(ob.AskRow, info)
		} else if info.Type == models.BID {
			ob.BidRow = append(ob.BidRow, info)
		}
	}
	return ob
}