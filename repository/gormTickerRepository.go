package repository

import (
	"context"
	models "tiffanyBlue/models"

	"github.com/jinzhu/gorm"
)

type gormTickerRepository struct {
	Conn *gorm.DB
}

// NewGormTickerRepository ...
func NewGormTickerRepository(Conn *gorm.DB) TickerRepository {
	return &gormTickerRepository{Conn}
}

func (g *gormTickerRepository) GetTickers(ctx context.Context) (ts []*models.Ticker, err error) {
	scope := g.Conn.Find(&ts)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return ts, scope.Error
}

func (g *gormTickerRepository) GetTicker(ctx context.Context, symbol string) (ticker *models.Ticker, err error) {
	scope := g.Conn.New()
	scope.Where(models.Ticker{TokenSymbol: symbol}).First(&ticker)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return ticker, scope.Error
}
