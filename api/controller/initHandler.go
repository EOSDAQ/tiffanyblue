package controller

import (
	"net/http"
	"time"

	mw "tiffanyBlue/api/middleware"
	"tiffanyBlue/conf"
	_Repo "tiffanyBlue/repository"
	"tiffanyBlue/service"
	"tiffanyBlue/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type (
	// TiffanyBlueStatus for common response status
	TiffanyBlueStatus struct {
		TRID       string      `json:"trID"`
		ResultCode string      `json:"resultCode"`
		ResultMsg  string      `json:"resultMsg"`
		ResultData interface{} `json:"resultData"`
	}
)

var mlog *zap.SugaredLogger

func init() {
	mlog, _ = util.InitLog("controller", "console")
}

// InitHandler ...
func InitHandler(tiffanyBlue conf.ViperConfig, e *echo.Echo, db *gorm.DB) (err error) {

	mlog, _ = util.InitLog("controller", tiffanyBlue.GetString("logmode"))
	timeout := time.Duration(tiffanyBlue.GetInt("timeout")) * time.Second

	// Default Group
	chart := e.Group("/")
	chart.File("/swagger.json", "swagger.json")
	chart.Use(mw.TransID())

	chartRepo := _Repo.NewGormChartRepository(db)
	chartSvc := service.NewChartService(chartRepo, timeout)
	newChartHTTPHandler(chart, chartSvc)

	api := e.Group("/api")
	ver := api.Group("/v1")
	sys := ver.Group("/eosdaq")
	sys.Use(mw.TransID())

	orderBook := sys.Group("/orderbook")
	ticker := sys.Group("/ticker")
	orderBookRepo := _Repo.NewGormOrderBookRepository(db)
	tokenRepo := _Repo.NewGormTokenRepository(db)
	orderBookSvc := service.NewOrderBookService(orderBookRepo, tokenRepo, timeout)
	tickerSvc := service.NewTickerService(tokenRepo, timeout)

	newOrderBookHTTPHandler(orderBook, orderBookSvc)
	newTickerHTTPHandler(ticker, tickerSvc)

	return nil
}

// HTTPChartHandler ...
type HTTPChartHandler struct {
	ChartService service.ChartService
}

func newChartHTTPHandler(eg *echo.Group, cs service.ChartService) {
	handler := &HTTPChartHandler{
		ChartService: cs,
	}

	eg.GET("", func(c echo.Context) error { return c.String(http.StatusOK, "tiffanyBlue API Alive!\n") })
	eg.POST("", func(c echo.Context) error { return c.String(http.StatusOK, "tiffanyBlue API Alive!\n") })

	eg.GET("config", handler.Config)
	eg.GET("symbol_info", handler.SymbolInfo)
	eg.GET("symbols", handler.Symbols)
	eg.GET("search", handler.Search)
	eg.GET("history", handler.History)
	eg.GET("marks", handler.Marks)
	eg.GET("timescale_marks", handler.TimeScale)
	eg.GET("time", handler.Time)
	eg.GET("quotes", handler.Quotes)
}

// HTTPOrderBookHandler ...
type HTTPOrderBookHandler struct {
	OrderBookService service.OrderBookService
}

func newOrderBookHTTPHandler(eg *echo.Group, obs service.OrderBookService) {
	handler := &HTTPOrderBookHandler{
		OrderBookService: obs,
	}

	eg.GET("/:symbol", handler.OrderBook)
}

// HTTPTickerHandler ...
type HTTPTickerHandler struct {
	TickerService service.TickerService
}

func newTickerHTTPHandler(eg *echo.Group, ts service.TickerService) {
	handler := &HTTPTickerHandler{
		TickerService: ts,
	}

	eg.GET("", handler.TickerList)
	eg.GET("/:symbol", handler.Ticker)
}
