# gobase Golang基础库

实际开发过程中，会写很多**八股文**代码，比如各种初始化

- 数据库
- Redis
- Elasticsearch
- 日志
- Http客户端
- 数据库Session操作
- 分布式ID生成器

本代码库就是集合这些八股文代码，在基本Google Wire的基础上，提供开箱即用的功能

## 使用方法

### 前提

所有的客户端都只有一个，比如

- Redis
- Xorm
- Elasticsearch

如果要用多个客户端，由于受Google Wire的限制，应该使用代码包装一层

```go
type DB1 gox.DatabaseConfig
type Engine1 *xorm.Engine

type DB2 gox.DatabaseConfig
type Engine2 *xorm.Engine

func NewDB1(config DB1) Engine1 {
	return Engine1(gobase.NewXorm(gox.DatabaseConfig(config)))
}

func NewDB2(config DB2) Engine2 {
    return Engine2(gobase.NewXorm(gox.DatabaseConfig(config)))
}
```

**如果发现出错上述类似的情况，理论上应该是设计上有问题，应该考虑设计规避**

### 关于App的问题

App是一个不得不提的问题，并**不建议使用**，原因是

- 提倡更明确的代码使用方式：你要使用什么功能，一定要明确的写出来
- 但是为了兼顾代码量，因为生成Id和记录日志通常是会大量使用的，所以没办法增加了这么一个工具类

### 依赖注入

如果使用Google Wire，那么方法特别简单，在任何Google Wire的Provider声明处增加

```go
func initApplication(path string) (application, error) {
	panic(wire.Build(
		// 顺序无关
		gobase.ProviderBaseSet,
		// 其它Provider
	))

	return application{}, nil
}
```

或者

```go
var ProviderXXXSet = wire.NewSet(
	// 顺序无关
	gobase.ProviderBaseSet,
	// 其它Provider
)
```

### 服务器容器

容器是拿来组织服务器的，写得很基础，之所以要说，是希望隐藏一些实现细节，真正的业务代码只需要关注我有几个服务器需要执行就好了

```go
package main

import (
	`embed`
	`io/fs`

	`live/rpc`

	`live/server`

	`github.com/storezhang/gobase`
)

//go:embed db/migration
var migration embed.FS

type application struct {
	container gobase.Container
}

func newApplication(container gobase.Container, server *server.Server, rpc *rpc.Rpc) (app application, err error) {
	container.AddServer(server)
	container.AddServer(rpc)

	// go embed会带上目录的所有层级，必须将最外层目录全部剥离
	var migrations fs.FS
	if migrations, err = fs.Sub(migration, "db/migration"); nil != err {
		return
	}
	container.AddMigration(migrations)

	app = application{
		container: container,
	}

	return
}

func (a application) run() (err error) {
	return a.container.Serve()
}
```

## 注意

使用过程中的注意点，如果不注意，有可能会出现No Provider的错误，如果不使用Google Wire，请使用各种NewXXX方法来代替初始化，强烈建议使用Google Wire来简化代码的组织结构

### 数据库

需要提供一个Debug的字段来确认是否开启调试模式

```go
// Debug 是否处于调试模式
type Debug bool
```

提供方式

```go
// 暴露数据库和程序整体配置
wire.FieldsOf(new(Config), "Debug"),
```

具体操作方式可以参考Google Wire的相关文档
