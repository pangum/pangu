package pangu

import (
	"github.com/goexl/simaqian"
)

var (
	_        = SimaqianLogger
	_ option = (*optionLogger)(nil)
)

type optionLogger struct {
	logger simaqian.Logger
}

// SimaqianLogger 日志
func SimaqianLogger(logger simaqian.Logger) *optionLogger {
	return &optionLogger{
		logger: logger,
	}
}

func (l *optionLogger) apply(options *options) {
	options.logger = l.logger
}
