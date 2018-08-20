package repository

import (
	"context"
	"fmt"
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

func (g *gormOrderBookRepository) GetOrderInfos(ctx context.Context, contract string) (obs []*models.OrderInfo, err error) {
	scope := g.Conn.Table(fmt.Sprintf("%s_order_books", contract)).Select("price, sum(volume) as volume, type").Group("price, type").Find(&obs)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return obs, scope.Error
}
