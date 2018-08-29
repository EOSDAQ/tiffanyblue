//go:generate swagger generate spec
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	ct "tiffanyBlue/api/controller"
	conf "tiffanyBlue/conf"
	_Repo "tiffanyBlue/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	svrInfo = fmt.Sprintf("tiffanyBlue %s(%s)", Version, BuildDate)
)

func init() {
	// use all cpu
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	TiffanyBlue := conf.TiffanyBlue
	if TiffanyBlue.GetBool("v") {
		fmt.Printf("%s\n", svrInfo)
		os.Exit(0)
	}
	TiffanyBlue.SetProfile()

	//f := apiLogFile("./tiffanyBlue-api.log")
	//defer f.Close()
	e := echoInit(TiffanyBlue, nil)

	// Prepare Server
	db := _Repo.InitDB(TiffanyBlue)
	defer db.Close()
	if err := ct.InitHandler(TiffanyBlue, e, db); err != nil {
		fmt.Println("InitHandler error : ", err)
		os.Exit(1)
	}

	startServer(TiffanyBlue, e)
}

func apiLogFile(logfile string) *os.File {
	// API Logging
	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("apiLogFile error : ", err)
		os.Exit(1)
	}
	return f
}

func echoInit(tiffanyBlue *conf.ViperConfig, apiLogF *os.File) (e *echo.Echo) {

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
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "tiffanyBlue API Alive!\n") })
	e.POST("/", func(c echo.Context) error { return c.String(http.StatusOK, "tiffanyBlue API Alive!\n") })

	loggerConfig := middleware.DefaultLoggerConfig
	//loggerConfig.Output = apiLogF

	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Logger.SetOutput(bufio.NewWriterSize(apiLogF, 1024*16))
	e.Logger.SetLevel(tiffanyBlue.APILogLevel())
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

func startServer(tiffanyBlue *conf.ViperConfig, e *echo.Echo) {
	// Start Server
	apiServer := fmt.Sprintf("0.0.0.0:%d", tiffanyBlue.GetInt("port"))
	log.Printf("%s => Starting server listen %s\n", svrInfo, apiServer)
	fmt.Printf(banner, svrInfo, apiServer)

	if err := e.Start(apiServer); err != nil {
		fmt.Println(err)
		e.Logger.Error(err)
	}
}
