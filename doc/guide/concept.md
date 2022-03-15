# 概念

本文主要介绍`盘古框架`的核心概念术语，有助于在使用中更准确的使用框架完成应用程序的快速开发

## 依赖

依赖注入是`盘古框架`的核心，依赖是指调用者依赖系统的组件（可以是结构体、方法或者变量），如果系统中有这个依赖，那么`盘古框架`就会提供正确的组件给调用方，使用方法

<<< @/../example/bootstrap.go#snippet{22}

## 注入

依赖注入是`盘古框架`的核心，依赖是指向`盘古框架`提供可以让别人使用的依赖组件（可以是结构体、方法或者变量），使用方法

<<< @/../example/conf/pangu.go#snippet{8}

## 服务

泛指一个可以长期执行的服务，可以是任何形式的服务，比如

- Http/RESTFul服务器
- gRPC服务器
- RocketMQ消费者

<<< @/../example/rest/server.go

## 命令

泛指一个可以被执行的命令，该命令可以在命令行中使用，比如

- 打印版本号
- 数据迁移
- 帮助信息
- 启动服务

<<< @/../example/command/test.go

## 参数

应用程序参数，命令行中使用，比如

- 指定配置文件路径
- 指定数据迁移脚本路径
- 指定应用程序内所需要的参数