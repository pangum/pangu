---

# @formatter:off
home: true
heroImage: /hero.png
heroText: Pangu
tagline: Golang应用快速开发框架
actionText: 快速上手 →
actionLink: /guide/
features:
  - title: 开箱即用
    details: 内置快速开发框架，引入模块即可快速开发自己的应用
  - title: 丰富的内置特性
    details: 内置支持版本输出、数据迁移、服务以及命令
  - title: 强大的扩展性
    details: 可以很方便的扩展整个框架而不侵入代码
  - title: 易于配置
    details: 框架的配置变得很容易，甚至可以无感知的配置系统
  - title: 线程安全
    details: 应用内单例，可以很方便的任何地方获得整个框架入口并作出配置而不用担心线程安全
  - title: 语义化
    details: 不需要记太多命令和参数以及配置项，语义化设计
footer: MIT Licensed | Copyright © 2021-present Storezhang
# @formatter:on
---

# 简单得不能再简单

```go
// @formatter:off
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
// @formatter:on
```

::: warning 警告 请使用Golang 1.16以上版本
:::
