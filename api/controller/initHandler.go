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
	"github.com/labstack/echo/middleware"
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
func InitHandler(tiffanyBlue *conf.ViperConfig, e *echo.Echo, db *gorm.DB) (err error) {

	mlog, _ = util.InitLog("controller", tiffanyBlue.GetString("env"))
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
	orderBookRepo := _Repo.NewGormOrderBookRepository(db)
	orderBookSvc := service.NewOrderBookService(orderBookRepo, timeout)
	newOrderBookHTTPHandler(orderBook, orderBookSvc)

	ticker := sys.Group("/ticker")
	tx := sys.Group("/tx")
	txRepo := _Repo.NewGormEosdaqTxRepository(db)
	txSvc := service.NewEosdaqTxService(txRepo, timeout)
	newEosdaqTxHTTPHandler(ticker, tx, txSvc)

	user := sys.Group("/user")
	userRepo := _Repo.NewGormUserRepository(db)
	userSvc := service.NewUserService(txRepo, orderBookRepo, timeout)
	newUserHTTPHandler(user, userSvc, burgundy.GetString("jwt_access_key"))

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

// HTTPEosdaqTxHandler ...
type HTTPEosdaqTxHandler struct {
	txService service.EosdaqTxService
}

func newEosdaqTxHTTPHandler(ticker *echo.Group, tx *echo.Group, txSvc service.EosdaqTxService) {
	handler := &HTTPEosdaqTxHandler{
		EosdaqTxService: txSvc,
	}

	ticker.GET("", handler.TickerList)
	ticker.GET("/:symbol", handler.Ticker)

	tx.GET("/:symbol", handler.SymbolTxList)
}

// HTTPUserHandler ...
type HTTPUserHandler struct {
	userService service.UserService
}

func newUserHTTPHandler(eg *echo.Group, user service.UserService, jwtkey string) {
	handler := &HTTPUserHandler{
		UserService: user,
	}

	if jwtkey != "" {
		eg.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey:  []byte(jwtkey),
			TokenLookup: "header:Authorization",
		}))
	}

	eg.GET("/:accountName", handler.UserTxList)
	eg.GET("/:accountName/symbol/:symbol", handler.UserSymbolTxList)
	eg.GET("/:accountName/orderbook", handler.UserOrderBook)
	eg.GET("/:accountName/symbol/:symbol/orderbook", handler.UserSymbolOrderBook)
}
