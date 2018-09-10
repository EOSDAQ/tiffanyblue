package service

import (
	"context"
	"tiffanyblue/models"
	"tiffanyblue/repository"
	"time"
)

type chartUsecase struct {
	chartRepo  repository.ChartRepository
	ctxTimeout time.Duration
}

// NewChartService ...
func NewChartService(cr repository.ChartRepository,
	timeout time.Duration) ChartService {
	return &chartUsecase{
		chartRepo:  cr,
		ctxTimeout: timeout,
	}
}

// GetByID ...
func (cus chartUsecase) GetByID(ctx context.Context, id string) (ct *models.Chart, err error) {
	innerCtx, cancel := context.WithTimeout(ctx, cus.ctxTimeout)
	defer cancel()

	return cus.chartRepo.GetByID(innerCtx, id)
}
