# 内置插件

为了方便开发者，`盘古`内置了一些扩展插件，方便快速开发的需求

## Http客户端

Http客户端是所有开发者经常使用的功能，所以选择内置在于`盘古`
中。Http客户端最终选择了[Resty](https://github.com/go-resty/resty)，相关操作完全可以通过查看Resty的文档来获得，同时在Resty的基础了做了一些简单的封闭，比如

- 良好的配置
- 代理
- 超时
- SSL证书
- 通用头
- 通过查询参数
- 通用Cookie
- 授权

具体配置可以参考[HttpClientConfig](https://github.com/storezhang/gox/blob/master/http_client.go)

## 日志

日志是绝大部分应用都需要支持的功能，所以`盘古`选择接入一个通用的日志库来解决此问题，[Glog](https://github.com/storezhang/glog)，该日志库支持以下几种日志库

- Zap Log
- Logrus
- Zero Log
- 系统内置日志

`盘古`只是接入了日志系统，但并没有对日志系统做任何修改，具体日志功能请参考官方文档
