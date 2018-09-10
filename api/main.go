//go:generate swagger generate spec
package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	ct "tiffanyblue/api/controller"
	mw "tiffanyblue/api/middleware"
	conf "tiffanyblue/conf"
	_Repo "tiffanyblue/repository"
	"tiffanyblue/util"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
)

const (
	banner = `
 _________  ___  ________ ________ ________  ________       ___    ___ ________  ___       ___  ___  _______      
|\___   ___\\  \|\  _____\\  _____\\   __  \|\   ___  \    |\  \  /  /|\   __  \|\  \     |\  \|\  \|\  ___ \     
\|___ \  \_\ \  \ \  \__/\ \  \__/\ \  \|\  \ \  \\ \  \   \ \  \/  / | \  \|\ /\ \  \    \ \  \\\  \ \   __/|    
     \ \  \ \ \  \ \   __\\ \   __\\ \   __  \ \  \\ \  \   \ \    / / \ \   __  \ \  \    \ \  \\\  \ \  \_|/__  
      \ \  \ \ \  \ \  \_| \ \  \_| \ \  \ \  \ \  \\ \  \   \/  /  /   \ \  \|\  \ \  \____\ \  \\\  \ \  \_|\ \ 
       \ \__\ \ \__\ \__\   \ \__\   \ \__\ \__\ \__\\ \__\__/  / /      \ \_______\ \_______\ \_______\ \_______\
        \|__|  \|__|\|__|    \|__|    \|__|\|__|\|__| \|__|\___/ /        \|_______|\|_______|\|_______|\|_______|
                                                          \|___|/                                                 
%s
 => Starting listen %s
`
)

var (
	// BuildDate for Program BuildDate
	BuildDate string
	// Version for Program Version
	Version string
	svrInfo = fmt.Sprintf("tiffanyblue %s(%s)", Version, BuildDate)
	mlog    *zap.SugaredLogger
)

func init() {
	// use all cpu
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	TiffanyBlue := conf.TiffanyBlue
	if TiffanyBlue.GetBool("version") {
		fmt.Printf("%s\n", svrInfo)
		os.Exit(0)
	}
	TiffanyBlue.SetProfile()
	mlog, _ = util.InitLog("main", TiffanyBlue.GetString("loglevel"))

	e := echoInit(TiffanyBlue)

	// Prepare Server
	db := _Repo.InitDB(TiffanyBlue)
	defer db.Close()
	if err := ct.InitHandler(TiffanyBlue, e, db); err != nil {
		mlog.Errorw("InitHandler", "err", err)
		os.Exit(1)
	}

	startServer(TiffanyBlue, e)
}

func echoInit(tiffanyblue *conf.ViperConfig) (e *echo.Echo) {

	// Echo instance
	e = echo.New()
	e.Debug = true

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
	}))

	// Ping Check
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "tiffanyblue API Alive!\n") })
	e.POST("/", func(c echo.Context) error { return c.String(http.StatusOK, "tiffanyblue API Alive!\n") })

	e.Use(mw.ZapLogger(mlog))
	e.HideBanner = true

	sigInit(e)

	return e
}

func sigInit(e *echo.Echo) chan os.Signal {

	// Signal
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		sig := <-sc
		e.Logger.Error("Got signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Error(err)
		}
		signal.Stop(sc)
		close(sc)
	}()

	return sc
}

func startServer(tiffanyblue *conf.ViperConfig, e *echo.Echo) {
	// Start Server
	apiServer := fmt.Sprintf("0.0.0.0:%d", tiffanyblue.GetInt("port"))
	mlog.Infow("Starting server", "info", svrInfo, "listen", apiServer)
	fmt.Printf(banner, svrInfo, apiServer)

	if err := e.Start(apiServer); err != nil {
		mlog.Errorw("End server", "err", err)
	}
}
