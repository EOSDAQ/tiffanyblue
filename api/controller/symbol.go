package controller

import (
	"context"
	"net/http"
	"tiffanyblue/models"

	"github.com/labstack/echo"
)

// TickerResponse ...
type TickerResponse struct {
	Tickers []*models.Token
}

// TickerList ...
func (h *HTTPSymbolHandler) TickerList(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	mlog.Debugw("TickerList", "tr", trID)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ts, err := h.sbService.GetTickers(ctx)
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
func (h *HTTPSymbolHandler) Ticker(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	symbol := c.Param("symbol")
	mlog.Debugw("ticker", "tr", trID, "symbol", symbol)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ticker, err := h.sbService.GetTicker(ctx, symbol)
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

// SymbolTxList ...
func (h *HTTPSymbolHandler) SymbolTxList(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	symbol := c.Param("symbol")
	mlog.Debugw("SymbolTxList", "tr", trID, "symbol", symbol)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	txs, err := h.sbService.GetSymbolTxList(ctx, symbol)
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
		ResultData: txs,
	})
}

// SymbolOrderBook ...
func (h *HTTPSymbolHandler) SymbolOrderBook(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	symbol := c.Param("symbol")
	mlog.Debugw("SymbolOrderBook", "tr", trID, "symbol", symbol)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	obs, err := h.sbService.GetSymbolOrderBook(ctx, symbol)
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
		ResultData: obs,
	})
}

// SymbolUserTxList ...
func (h *HTTPSymbolHandler) SymbolUserTxList(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	symbol := c.Param("symbol")
	accountName := c.Param("accountName")
	mlog.Debugw("SymbolUserTxList", "tr", trID, "symbol", symbol, "accountName", accountName)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	txs, err := h.userService.GetUserSymbolTxList(ctx, accountName, symbol)
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
		ResultData: txs,
	})
}

// SymbolUserOrderInfos ...
func (h *HTTPSymbolHandler) SymbolUserOrderInfos(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	symbol := c.Param("symbol")
	accountName := c.Param("accountName")
	mlog.Debugw("SymbolUserOrderInfos", "tr", trID, "symbol", symbol, "accountName", accountName)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	obinfos, err := h.userService.GetUserSymbolOrderInfos(ctx, accountName, symbol)
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
		ResultData: obinfos,
	})
}
