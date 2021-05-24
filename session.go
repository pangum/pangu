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
	engine *xorm.Engine
	logger glog.Logger
}

// NewSession 事务控制
func NewSession(engine *xorm.Engine, logger glog.Logger) Session {
	return Session{
		engine: engine,
		logger: logger,
	}
}

func (s *Session) Tx(fun func(tx *xorm.Session) error, fields ...gox.Field) (err error) {
	tx := s.engine.NewSession()
	defer s.close(tx, fields...)

	if err = fun(tx); nil != err {
		s.rollback(tx, fields...)
	} else {
		s.commit(tx, fields...)
	}

	return
}

func (s *Session) commit(tx *xorm.Session, fields ...gox.Field) {
	if err := tx.Commit(); nil != err {
		fun, _, line, _ := runtime.Caller(1)

		logFields := make([]gox.Field, 0, len(fields)+4)
		logFields = append(logFields, field.String("fun", runtime.FuncForPC(fun).Name()))
		logFields = append(logFields, field.Int("line", line))
		logFields = append(logFields, fields...)
		logFields = append(logFields, field.Error(err))
		s.logger.Error("提交数据库事务出错", logFields...)
	}
}

func (s *Session) close(tx *xorm.Session, fields ...gox.Field) {
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

func (s *Session) rollback(tx *xorm.Session, fields ...gox.Field) {
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
