package repository

import (
	"context"
	models "tiffanyBlue/models"

	"github.com/jinzhu/gorm"
)

type gormOrderBookRepository struct {
	Conn *gorm.DB
}

// NewGormOrderBookRepository ...
func NewGormOrderBookRepository(Conn *gorm.DB) OrderBookRepository {
	return &gormOrderBookRepository{Conn}
}

// GetOrderInfos ...
func (g *gormOrderBookRepository) GetOrderInfos(ctx context.Context, symbol string) (obs []*models.OrderInfo, err error) {
	scope := g.Conn.Table("order_books").
		Select("price, sum(volume) as volume, type").
		Where("order_symbol = ?", symbol).
		Group("price, type").Find(&obs)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return obs, scope.Error
}

// GetUserOrderInfos ...
func (g *gormOrderBookRepository) GetUserOrderInfos(ctx context.Context, accountName string) (obs []*models.OrderInfo, err error) {
	scope := g.Conn.Table("order_books").
		Select("order_symbol, order_time, price, volume, type").
		Where("account_name = ?", accountName).
		Find(&obs)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return obs, scope.Error
}

// GetUserSymbolOrderInfos ...
func (g *gormOrderBookRepository) GetUserSymbolOrderInfos(ctx context.Context, accountName, symbol string) (obs []*models.OrderInfo, err error) {
	scope := g.Conn.Table("order_books").
		Select("order_symbol, order_time, price, volume, type").
		Where("account_name = ? and order_symbol = ?", accountName, symbol).
		Find(&obs)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return obs, scope.Error
}
