# 命令

可以这样理解一个命令
> 一个短时间内运行的功能

比如说

- 帮助
- 数据迁移
- 打印版本信息
- 其它应用程序所需要的命令

## 定义一个命令

在`盘古`里面很方便定义一个服务，只需要实现`app.Command`接口就可以了

<<< @/../app/command.go

## 命令实例

例子代码里面实现了一个简单的RESTFul服务，要实现复杂一点的服务，只需要在这上面增加即可

<<< @/../example/rest/server.go

建议

- 分包`controller`和`handler`，这在Java或者其它语言的Web开发中比较常见
- 使用[Echox](https://github.com/storezhang/echox)来开发RESTFul服务器
- 增加一个`pangu.go`来组织依赖
