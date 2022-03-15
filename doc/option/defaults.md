# 默认值

配置`盘古`是否处理配置的默认值（应用程序配置，默认开启本选项）

::: warning 注意

`盘古`框架默认开启了默认值，默认值也可以使用环境变量，具体信息参看[Mengpo使用教程](https://github.com/goexl/mengpo)
:::

## 关闭

<<< @/../example/option/defaults_disable.go

## 高级使用

<<< @/../example/option/defaults_advanced.go

在里面有两个变量，分别表示

- `Folder`：如果环境变量`PLUGIN_FOLDER`有值，则使用，否则取环境变量`FOLDER`
- `Scripts`：如果环境变量`PLUGIN_SCRIPT`有值，则使用，否则取环境变量`SCRIPT`，如果环境变量`SCRIPT`也没有被设置，则使用默认值`build`

还有更多的使用方式，大家可以参考文档[Mengpo使用教程](https://github.com/goexl/mengpo)来了解

