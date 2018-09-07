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

	symbol := sys.Group("/symbol")
	obRepo := _Repo.NewGormOrderBookRepository(db)
	txRepo := _Repo.NewGormEosdaqTxRepository(db)
	sbSvc := service.NewSymbolService(obRepo, txRepo, timeout)
	userSvc := service.NewUserService(obRepo, txRepo, timeout)
	newSymbolHTTPHandler(symbol, sbSvc, userSvc, tiffanyBlue.GetString("jwt_access_key"))

	user := sys.Group("/user")
	newUserHTTPHandler(user, userSvc, tiffanyBlue.GetString("jwt_access_key"))

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

// HTTPSymbolHandler ...
type HTTPSymbolHandler struct {
	sbService   service.SymbolService
	userService service.UserService
}

func newSymbolHTTPHandler(eg *echo.Group, sbSvc service.SymbolService, userSvc service.UserService, jwtkey string) {
	handler := &HTTPSymbolHandler{
		sbService:   sbSvc,
		userService: userSvc,
	}

	eg.GET("", handler.TickerList)
	eg.GET("/:symbol", handler.Ticker)
	eg.GET("/:symbol/tx", handler.SymbolTxList)
	eg.GET("/:symbol/orderbook", handler.SymbolOrderBook)

	user := eg.Group("/:symbol/user")
	if jwtkey != "" {
		user.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey:  []byte(jwtkey),
			TokenLookup: "header:Authorization",
		}))
	}

	user.GET("/:accountName/tx", handler.SymbolUserTxList)
	user.GET("/:accountName/orderbook", handler.SymbolUserOrderInfos)
}

// HTTPUserHandler ...
type HTTPUserHandler struct {
	userService service.UserService
}

func newUserHTTPHandler(eg *echo.Group, user service.UserService, jwtkey string) {
	handler := &HTTPUserHandler{
		userService: user,
	}

	if jwtkey != "" {
		eg.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey:  []byte(jwtkey),
			TokenLookup: "header:Authorization",
		}))
	}

	eg.GET("/:accountName/tx", handler.UserTxList)
	eg.GET("/:accountName/orderbook", handler.UserOrderInfos)
}
