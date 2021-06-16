package pangu

import (
	`github.com/storezhang/glog`
)

var _ option = (*optionLogger)(nil)

type optionLogger struct {
	logger glog.Logger
}

// GLogger 配置日志
func GLogger(logger glog.Logger) *optionLogger {
	return &optionLogger{
		logger: logger,
	}
}

func (l *optionLogger) apply(options *options) {
	options.logger = l.logger
}
