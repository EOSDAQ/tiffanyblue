package controller

import (
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
		TRID       string `json:"trID"`
		ResultCode string `json:"resultCode"`
		ResultMsg  string `json:"resultMsg"`
		ResultData string `json:"resultData"`
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
