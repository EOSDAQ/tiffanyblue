package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

// UserTxList ...
func (h *HTTPUserHandler) UserTxList(c echo.Context) (err error) {

	type request struct {
		Page uint `query:"page"`
	}

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	accountName := c.Param("accountName")
	req := &request{}

	if err = c.Bind(req); err != nil || accountName == "" {
		mlog.Errorw("UserTxList bind error ", "trID", trID, "req", *req, "err", err)
		return c.JSON(http.StatusBadRequest, TiffanyBlueStatus{
			TRID:       trID,
			ResultCode: "1101",
			ResultMsg:  "Invalid Parameter",
		})
	}

	if req.Page == 0 {
		req.Page = 1
	}
	mlog.Debugw("UserTxList", "tr", trID, "accountName", accountName, "req", *req)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	txs, err := h.userService.GetUserTxList(ctx, accountName, req.Page)
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

// UserOrderInfos ...
func (h *HTTPUserHandler) UserOrderInfos(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	accountName := c.Param("accountName")

	mlog.Debugw("UserOrderBook", "tr", trID, "user", accountName)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	obinfos, err := h.userService.GetUserOrderInfos(ctx, accountName)
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
