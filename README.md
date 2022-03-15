# 盘古 Golang应用程序框架

一个Golang应用程序快速开发框架，有以下特性

- 快速开发
- 内置配置文件加载
    - `Yaml`
    - `Toml`
    - `Json`
    - `XML`
- 内置强大的数据验证
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

`Pangu`使用非常简单，只需要定义两部分代码`启动方法`和`启动器`

启动方法代码如下

```go
package main

import (
    `github.com/pangum/pangu`
    `github.com/pangum/pangum`
)

func main() {
    panic(pangu.New(
        pangu.Name(`ziyunix`),
        pangu.Banner(`Ziyunix Server`, pangu.BannerTypeAscii),
    ).Run(newBootstrap))
}
```

启动器的代码如下

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

> `Pangu`有非常多的配置项，请参看[**使用文档**](https://pangu.pangum.tech)

## 文档

[点击这里查看最新文档](https://pangu.pangum.tech)

## 使用示例

[点击这里查看最新的代码示例](example)

## 交流

![微信群](doc/.vuepress/public/communication/wxwork.jpg)

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)
