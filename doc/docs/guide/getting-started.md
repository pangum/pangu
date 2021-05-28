# 快速下手

::: warning 前提条件
<!--@formatter:off-->
Golang需要1.16及以上版本
<!--@formatter:on-->
:::

本文会帮助你从头搭建一个简单的Golang应用

1. 创建并进入一个新目录

``` bash
mkdir pangu-app && cd pangu-app
```

2. 初始化Go Module

``` bash
go mod init
```

3. 添加pangu作为依赖

``` go
module server

go 1.16

require (
	github.com/rs/xid v1.3.0
	github.com/storezhang/echox/v2 v2.0.3
	github.com/storezhang/glog v1.0.5
	github.com/storezhang/gox v1.4.10
	github.com/storezhang/pangu v1.2.0
	github.com/storezhang/pangu-database v1.0.2
	github.com/storezhang/una v1.0.7
	github.com/storezhang/uoa v1.1.0
	xorm.io/builder v0.3.9
	xorm.io/xorm v1.1.0
)
```

4. 创建启动器

``` go
package main

import (
	`embed`
	`io/fs`

	`server/conf`
	`server/repository`
	`server/rest`
	`server/service`

	`github.com/storezhang/pangu`
)

type bootstrap struct {
	application *pangu.Application
}

func newBootstrap(application *pangu.Application) pangu.Bootstrap {
	return &bootstrap{
		application: application,
	}
}

func (b *bootstrap) Setup() (err error) {
	if err = b.provides(); nil != err {
		return
	}

	if err = b.application.Invoke(func(server *rest.Server) error {
		return b.application.AddServes(server)
	}); nil != err {
		return
	}

	return
}

func (b *bootstrap) provides() (err error) {
	// 顺序无关
	if err = conf.Provides(b.application); nil != err {
		return
	}
	if err = rest.Provides(b.application); nil != err {
		return
	}
	if err = repository.Provides(b.application); nil != err {
		return
	}
	if err = service.Provides(b.application); nil != err {
		return
	}

	return
}
```

5. 创建启动文件

``` go
package main

import (
	`github.com/storezhang/pangu`
	_ `github.com/storezhang/pangu-database`
)

func main() {
	panic(pangu.New(
		pangu.Name("archtech"),
		pangu.Banner("archtech", pangu.BannerTypeAscii),
	).Run(newBootstrap))
}
```

6. 启动本运行服务器

``` bash
go build
```
