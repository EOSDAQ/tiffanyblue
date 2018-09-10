// Package repository ...
//
// Repository will store any Database handler.
// Querying, or Creating/ Inserting into any database will stored here.
// This layer will act for CRUD to database only.
// No business process happen here. Only plain function to Database.
//
// This layer also have responsibility to choose what DB will used in Application.
// Could be Mysql, MongoDB, MariaDB, Postgresql whatever, will decided here.
//
// If using ORM, this layer will control the input, and give it directly to ORM services.
//
// If calling microservices, will handled here. Create HTTP Request to other services, and sanitize the data.
// This layer, must fully act as a repository. Handle all data input - output no specific logic happen.
//
// This Repository layer will depends to Connected DB , or other microservices if exists.
package repository

import (
	"context"
	"fmt"
	"os"

	"tiffanyblue/conf"
	models "tiffanyblue/models"
	"tiffanyblue/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql version
	"go.uber.org/zap"
)

var mlog *zap.SugaredLogger

func init() {
	mlog, _ = util.InitLog("repository", "console")
}

// InitDB ...
func InitDB(tiffanyblue *conf.ViperConfig) *gorm.DB {

	mlog, _ = util.InitLog("repository", tiffanyblue.GetString("loglevel"))

	mlog.Debugw("InitDB ",
		"host",
		tiffanyblue.GetString("db_host"),
		"user",
		tiffanyblue.GetString("db_user"),
		"name",
		tiffanyblue.GetString("db_name"),
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		tiffanyblue.GetString("db_user"),
		tiffanyblue.GetString("db_pass"),
		tiffanyblue.GetString("db_host"),
		tiffanyblue.GetInt("db_port"),
		tiffanyblue.GetString("db_name"),
	)
	dbConn, err := gorm.Open("mysql", dbURI) //mysql version
	if err != nil {
		fmt.Println("InitDB", err)
		os.Exit(1)
	}
	dbConn.DB().SetMaxIdleConns(100)
	dbConn.LogMode(true)
	return dbConn
}

// ChartRepository ...
type ChartRepository interface {
	GetByID(ctx context.Context, id string) (*models.Chart, error)
}

// OrderBookRepository ...
type OrderBookRepository interface {
	GetOrderInfos(ctx context.Context, symbol string) (obs []*models.OrderInfo, err error)
	GetUserOrderInfos(ctx context.Context, accountName string) (obs []*models.UserOrderInfo, err error)
	GetUserSymbolOrderInfos(ctx context.Context, accountName, symbol string) (obs []*models.UserOrderInfo, err error)
}

// EosdaqTxRepository ...
type EosdaqTxRepository interface {
	GetTickers(ctx context.Context) (ts []*models.Token, err error)
	GetTicker(ctx context.Context, symbol string) (token *models.Token, err error)

	GetSymbolTxList(ctx context.Context, symbol string) (txs []*models.EosdaqTx, err error)
	GetUserTxList(ctx context.Context, accountName string, page uint) (txs []*models.EosdaqTx, err error)
	GetUserSymbolTxList(ctx context.Context, accountName, symbol string) (txs []*models.EosdaqTx, err error)
}
