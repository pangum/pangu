# 服务

可以这样理解一个服务
> 一个可以长期运行的功能

比如说

- Http服务
- gRPC服务
- Socket服务

## 定义一个服务

在`盘古`里面很方便定义一个服务，只需要实现`pangu.Serve`接口就可以了

<<< @/../app/serve.go

## 服务实例

例子代码里面实现了一个简单的RESTFul服务，要实现复杂一点的服务，只需要在这上面增加即可

<<< @/../example/rest/server.go

建议

- 分包`controller`和`handler`，这在Java或者其它语言的Web开发中比较常见
- 使用[Echox](https://github.com/storezhang/echox)来开发RESTFul服务器
- 增加一个`pangu.go`来组织依赖
