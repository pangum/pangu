package pangu

import (
	`runtime`

	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`xorm.io/xorm`
)

// Session 事务控制
type Session struct {
	logger glog.Logger
}

// NewSession 事务控制
func NewSession(logger glog.Logger) Session {
	return Session{
		logger: logger,
	}
}

// Close 关闭事务
func (s *Session) Close(tx *xorm.Session, fields ...gox.Field) {
	if err := tx.Close(); nil != err {
		fun, _, line, _ := runtime.Caller(1)

		logFields := make([]gox.Field, 0, len(fields)+4)
		logFields = append(logFields, field.String("fun", runtime.FuncForPC(fun).Name()))
		logFields = append(logFields, field.Int("line", line))
		logFields = append(logFields, fields...)
		logFields = append(logFields, field.Error(err))
		s.logger.Error("关闭数据库事务出错", logFields...)
	}
}

// Rollback 回退事务
func (s *Session) Rollback(tx *xorm.Session, fields ...gox.Field) {
	if err := tx.Rollback(); nil != err {
		fun, _, line, _ := runtime.Caller(1)

		logFields := make([]gox.Field, 0, len(fields)+4)
		logFields = append(logFields, field.String("fun", runtime.FuncForPC(fun).Name()))
		logFields = append(logFields, field.Int("line", line))
		logFields = append(logFields, fields...)
		logFields = append(logFields, field.Error(err))
		s.logger.Error("回退数据库事务出错", logFields...)
	}
}
