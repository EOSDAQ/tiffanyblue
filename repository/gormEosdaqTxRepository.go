package repository

import (
	"context"
	models "tiffanyblue/models"

	"github.com/jinzhu/gorm"
)

type gormEosdaqTxRepository struct {
	Conn  *gorm.DB
	limit uint
}

// NewGormEosdaqTxRepository ...
func NewGormEosdaqTxRepository(Conn *gorm.DB) EosdaqTxRepository {
	return &gormEosdaqTxRepository{Conn, 30}
}

func (g *gormEosdaqTxRepository) GetTickers(ctx context.Context) (ts []*models.Token, err error) {
	scope := g.Conn.Find(&ts)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return ts, scope.Error
}

func (g *gormEosdaqTxRepository) GetTicker(ctx context.Context, symbol string) (token *models.Token, err error) {
	token = &models.Token{}
	scope := g.Conn.Where("symbol = ?", symbol).First(&token)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return token, scope.Error
}

func (g *gormEosdaqTxRepository) GetSymbolTxList(ctx context.Context, symbol string) (txs []*models.EosdaqTx, err error) {
	scope := g.Conn.Where(&models.EosdaqTx{OrderSymbol: symbol}).
		Order("id desc").
		Limit(g.limit).
		Find(&txs)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return txs, scope.Error
}

func (g *gormEosdaqTxRepository) GetUserTxList(ctx context.Context, accountName string, page uint) (txs []*models.EosdaqTx, err error) {
	scope := g.Conn.Where(models.EosdaqTx{
		EOSData: &models.EOSData{AccountName: accountName},
	}).
		Order("id desc").
		Offset((page - 1) * g.limit).
		Limit(g.limit).
		Find(&txs)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return txs, scope.Error
}

func (g *gormEosdaqTxRepository) GetUserSymbolTxList(ctx context.Context, accountName, symbol string) (txs []*models.EosdaqTx, err error) {
	scope := g.Conn.Where(models.EosdaqTx{OrderSymbol: symbol, EOSData: &models.EOSData{AccountName: accountName}}).
		Order("id desc").
		Limit(g.limit).
		Find(&txs)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return txs, scope.Error
}
