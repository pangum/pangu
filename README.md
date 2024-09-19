# `盘古` Golang应用程序框架
[![编译状态](https://github.ruijc.com:20443/api/badges/pangum/pangu/status.svg)](https://github.ruijc.com:20443/pangum/pangu)
[![Golang质量](https://goreportcard.com/badge/github.com/pangum/pangu)](https://goreportcard.com/report/github.com/pangum/pangu)
![版本](https://img.shields.io/github/go-mod/go-version/pangum/pangu)
![仓库大小](https://img.shields.io/github/repo-size/pangum/pangu)
![最后提交](https://img.shields.io/github/last-commit/pangum/pangu)
![授权协议](https://img.shields.io/github/license/pangum/pangu)
![语言个数](https://img.shields.io/github/languages/count/pangum/pangu)
![最佳语言](https://img.shields.io/github/languages/top/pangum/pangu)
![星星个数](https://img.shields.io/github/stars/pangum/pangu?style=social)

一个Golang应用程序快速开发框架，有以下特性

- 快速开发
- 内置配置文件加载
  - `Json`/`Jsonc`/`Json5`
  - `XML`
  - 其它
    - `YAML`，通过引入`github.com/pangum/yaml`支持
    - `TOML`，通过引入`github.com/pangum/toml`支持
    - 携程`Apollo`，通过引入`github.com/pangum/apollo`支持
    - 阿里`Nacos`，通过引入`github.com/pangum/nacos`支持
  - 很方便定制自己的配置文件加载器
  - 零配置，可完全通过`环境变量`完成配置
  - 支持配置文件覆盖
  - 支持按`模块`划分配置文件，比如可以按如下模式划分配置文件
    - `application.json/jsonc/xml/toml/yaml/yml`，提供应用程序本身需要的相关配置
    - `logging.json/jsonc/xml/toml/yaml/yml`，提供日志相关配置
    - `database.json/jsonc/xml/toml/yaml/yml`，提供数据库相关配置
    - `grpc.json/jsonc/xml/toml/yaml/yml`，提供`gRPC`相关配置
    - `...`其它配置文件
    - 所有配置文件格式可以混用，比如配置可以是`json`、`jsonc`、`json5`、`yaml`、`yml`、`toml`等任意格式，只要正确的引入了相应的`加载器`
  - 支持个性化配置文件编写，对于任何一个配置字段，可以有如下形式（比如配置项是`logggingLevel`）
    - `原始配置项`即`loggingLevel`
    - `下划线形式`即`logging_level`
    - `中划线形式`即`logging-level`
- 内置强大的数据验证
  - 内置强大的且非常多的数据验证器
  - 很方便的定制自己的数据验证器
- 强大的扩展功能
- 强大的配置功能
  - 徽标
  - 帮助信息
  - 命令行
  - 一切可配置点皆可配置
- 线程安全
- 内置依赖注入
- 语义化
  - 方便使用
  - 多态行为，同一个调用在不同的场景下有不同的意义

## 快速开始

`Pangu`使用非常简单，只需要定义两部分代码`启动方法`和`启动器`以及`依赖关系`，对于`依赖关系`

- 可以没有依赖，如果没有，就不需要处理依赖
- 一般稍微大型一点的系统都会有复杂的依赖关系，而这正是`盘古`框架的强项
- 对于没有依赖或者依赖非常少的项目，使用`盘古`反而会使代码变多，这个需要大家酌情考虑

`启动方法`代码如下

```go
package main

import (
    `github.com/pangum/pangu`
    `github.com/pangum/pangum`
)

func main() {
    panic(pangu.New(
        pangu.Named(`ziyunix`),
        pangu.Banner(`Ziyunix Server`, pangu.BannerTypeAscii),
    ).Run(newBootstrap))
}
```

`启动器`的代码如下

```go
package main

type bootstrap struct {
    application *pangu.Application
}

func newBootstrap(application *pangu.Application) pangu.Bootstrap {
    return &bootstrap{
        application: application,
    }
}
```

`依赖关系`建议像如下处理

```go
package rest

import (
	`github.com/pangum/pangu`
)

type Server struct {}

func newServer(/* 如果有依赖，可以在这里增加依赖：api *Api */) *Server {
	return new(Server)
}

func init() {
	pangu.New().Dependencies(
		newServer,
		// 其它依赖
		// ...
	)
}
```

> `Pangu`有非常多的配置项，请参看[**使用文档**](https://pangu.pangum.tech)

## 文档

[点击这里查看最新文档](https://pangu.pangum.tech)

## 使用示例

[点击这里查看最新的代码示例](example)

## 项目实践

- [基础项目实战代码](https://github.com/pangum/example)
  - 数据库操作
  - 数据库迁移
  - RESTFul接口
  - 配置加载
  - MQTT操作

## 交流

![微信群](doc/.vuepress/public/communication/wxwork.jpg)

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

## 感谢Jetbrains

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢
[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg)](https://www.jetbrains.com/?from=pangum/pangu)
