package app

import (
	`github.com/goexl/gox`
	`github.com/goexl/simaqian`
)

// defaultService 默认应用服务
type defaultService struct {
	logger    simaqian.Logger
	snowflake *gox.Snowflake
}

// NewDefaultService 创建默认的应用服务
func NewDefaultService(logger simaqian.Logger, snowflake *gox.Snowflake) Service {
	return &defaultService{
		logger:    logger,
		snowflake: snowflake,
	}
}

func (ds *defaultService) NextId() int64 {
	return ds.snowflake.NextId()
}

func (ds *defaultService) Debug(msg string, fields []gox.Field) {
	ds.logger.Debug(msg, fields...)
}

func (ds *defaultService) Info(msg string, fields []gox.Field) {
	ds.logger.Info(msg, fields...)
}

func (ds *defaultService) Warn(msg string, fields []gox.Field) {
	ds.logger.Warn(msg, fields...)
}

func (ds *defaultService) Error(msg string, fields []gox.Field) {
	ds.logger.Error(msg, fields...)
}

func (ds *defaultService) Panic(msg string, fields []gox.Field) {
	ds.logger.Panic(msg, fields...)
}

func (ds *defaultService) Fatal(msg string, fields []gox.Field) {
	ds.logger.Fatal(msg, fields...)
}
