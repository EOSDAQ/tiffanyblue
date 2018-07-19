package models

import (
	"tiffanyBlue/util"

	"go.uber.org/zap"
)

var mlog *zap.SugaredLogger

func init() {
	mlog, _ = util.InitLog("models", "console")
}
