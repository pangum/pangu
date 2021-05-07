# pangu Golang应用程序框架

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

### 创建主方法main.go

```go
package main

import (
	`github.com/storezhang/pangu`
)

func main(){
	panic(pangu.New(
		pangu.Name("example"),
		pangu.Banner("./banner.txt", pangu.BannerTypeFile),
	).Run(newBootstrap))
}
```

### 创建程序启动器

```go
package main

import (
	`example/conf`
	`example/rest`

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
	if err = b.inject(); nil != err {
		return
	}
	if err = b.application.Get(func(server *rest.Server) error {
		return b.application.AddServe(server)
	}); nil != err {
		return
	}

	return
}

func (b *bootstrap) inject() (err error) {
	if err = conf.Provide(b.application); nil != err {
		return
	}
	if err = rest.Provide(b.application); nil != err {
		return
	}

	return
}

```

### 启动程序

```shell
./example serve
```

## 内置命令

### 配置文件 -c --conf --config --configuration

配置文件有默认值，顺序依次是

- ./conf/application.yaml
- ./conf/application.yml
- ./conf/application.toml
- ./conf/application.json
- ./application.yaml
- ./application.yml
- ./application.toml
- ./application.json
- 参数指定

### 输出版本号 v version V Version

### 提供服务 s serve S Serve
