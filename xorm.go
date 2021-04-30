package pangu

import (
	`strings`

	`github.com/storezhang/gox`
	`xorm.io/core`
	`xorm.io/xorm`
	`xorm.io/xorm/log`
)

// NewXormEngine 创建Xorm操作引擎
func NewXormEngine(config gox.DatabaseConfig, debug gox.Debug) (engine *xorm.Engine, err error) {
	if engine, err = xorm.NewEngine(config.Type, config.Dsn()); nil != err {
		return
	}

	// 调试模式下打开各种可调试的选项
	if debug {
		engine.ShowSQL(true)
		engine.Logger().SetLevel(log.LOG_DEBUG)
	}

	// 配置数据库连接池
	engine.SetMaxOpenConns(config.Connection.MaxOpen)
	engine.SetMaxIdleConns(config.Connection.MaxIdle)
	engine.SetConnMaxLifetime(config.Connection.MaxLifetime)

	// 测试数据库连接成功
	if err = engine.Ping(); nil != err {
		return
	}

	// 设置名称转换（列名及表名）
	core.NewCacheMapper(core.GonicMapper{})
	if "" != strings.TrimSpace(config.Prefix) {
		core.NewPrefixMapper(core.GonicMapper{}, config.Prefix)
	}
	if "" != strings.TrimSpace(config.Suffix) {
		core.NewSuffixMapper(core.GonicMapper{}, config.Suffix)
	}

	return
}
