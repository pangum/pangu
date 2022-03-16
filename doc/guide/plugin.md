# 插件

插件是指一类可以被`盘古`认识并使用的独立的功能封装框架，使用类似于`Springboot`自动配置的形式提供插件。 这样可以很方便开发者使用插件，完全不需要自己开发只需要使用（当然会有一引起配置项）。
`盘古`的插件具有以下特征

- 自动配置
- 内置丰富的插件
- 方便开发自己的插件

## 内置插件

为了方便开发者，`盘古`内置了一些扩展插件，方便快速开发的需求

- [Web开发](https://github.com/pangum/web)
- [Http客户端](https://github.com/pangum/http)
- [数据库](https://github.com/pangum/database)
- [Pulsar](https://github.com/pangum/pulsar)
- [Mqtt](https://github.com/pangum/mqtt)
- [Redis](https://github.com/pangum/redis)
- [gRPC](https://github.com/pangum/grpc)
- [日志](https://github.com/pangum/logging)

## 插件开发

`盘古`可以很方便的接入自己的插件，插件可以是

- 服务
- 命令
- 参数

具体如何开发插件[请移步](/plugin)
