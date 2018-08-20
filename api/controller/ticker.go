package controller

import (
	"context"
	"net/http"
	"tiffanyBlue/models"

	"github.com/labstack/echo"
)

// TickerResponse ...
type TickerResponse struct {
	Tickers []*models.Token
}

// Ticker ...
func (h *HTTPTickerHandler) TickerList(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	mlog.Infow("TickerList", "tr", trID)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ts, err := h.TickerService.GetTickers(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, TiffanyBlueStatus{
			ResultCode: "1000",
			ResultMsg:  err.Error(),
			TRID:       trID,
		})
	}

	return c.JSON(http.StatusOK, TiffanyBlueStatus{
		ResultCode: "0000",
		ResultMsg:  "OK",
		TRID:       trID,
		ResultData: TickerResponse{ts},
	})
}

// Ticker ...
func (h *HTTPTickerHandler) Ticker(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	symbol := c.Param("symbol")
	mlog.Infow("ticker", "tr", trID, "symbol", symbol)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ticker, err := h.TickerService.GetTicker(ctx, symbol)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, TiffanyBlueStatus{
			ResultCode: "1000",
			ResultMsg:  err.Error(),
			TRID:       trID,
		})
	}

	return c.JSON(http.StatusOK, TiffanyBlueStatus{
		ResultCode: "0000",
		ResultMsg:  "OK",
		TRID:       trID,
		ResultData: TickerResponse{[]*models.Token{ticker}},
	})
}
