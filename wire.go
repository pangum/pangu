package pangu

import (
	`github.com/google/wire`
	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
	`github.com/storezhang/pangu/command`
)

// ProviderSet 基础库
var ProviderSet = wire.NewSet(
	// 注入雪花生成器
	gox.NewSnowflake,
	// 注入日志
	glog.NewLogger,
	// 注入Http客户端
	NewResty,
	// 注入数据库
	NewXormEngine, NewSession,
	// 注入Redis
	NewRedis,
	// 注入Elasticsearch
	NewElasticsearch,
	// 注入版本信息
	appName, appVersion, buildVersion, buildTime, scmRevision, scmBranch, goVersion,
	// 注入命令
	command.NewServe, command.NewVersion,
	// 注入应用程序
	// newApp, NewBootstrap, NewApplication,
)
