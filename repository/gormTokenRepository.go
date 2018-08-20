package repository

import (
	"context"
	models "tiffanyBlue/models"

	"github.com/jinzhu/gorm"
)

type gormTokenRepository struct {
	Conn *gorm.DB
}

// NewGormTokenRepository ...
func NewGormTokenRepository(Conn *gorm.DB) TokenRepository {
	return &gormTokenRepository{Conn}
}

func (g *gormTokenRepository) GetTokens(ctx context.Context) (ts []*models.Token, err error) {
	scope := g.Conn.Find(&ts)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return ts, scope.Error
}

func (g *gormTokenRepository) GetToken(ctx context.Context, symbol string) (token *models.Token, err error) {
	scope := g.Conn.New()
	scope.Where(models.Token{Symbol: symbol}).First(&token)
	if scope.RowsAffected == 0 {
		return nil, nil
	}
	return token, scope.Error
}
