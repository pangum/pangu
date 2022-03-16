# 名称

可以配置应用程序显示帮助信息的名称，有三种方法可以定制应用的名称

- 编译时添加编译参数

```shell
go build -ldflags "-s -X 'github.com/pangum/pangu.Name=$NAME"
```

- 使用`Drone`的`盘古`插件

[参考文档](https://github.com/pangum/drone)

- 运行时使用配置选项

<<< @/../example/option/named.go
