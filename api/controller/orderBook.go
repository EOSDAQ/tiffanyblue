package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

// OrderBook ...
func (h *HTTPOrderBookHandler) OrderBook(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	symbol := c.Param("symbol")
	mlog.Debugw("orderbook", "tr", trID, "symbol", symbol)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	obs, err := h.OrderBookService.GetOrderBooks(ctx, symbol)
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
