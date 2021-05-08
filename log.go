package pangu

import (
	`github.com/storezhang/glog`
	`go.uber.org/zap`
)

func newZapLogger() (logger *glog.ZapLogger, err error) {
	var zapLogger *zap.Logger
	if zapLogger, err = zap.NewProduction(); nil != err {
		return
	}

	logger = glog.NewZapLogger(zapLogger)

	return
}
