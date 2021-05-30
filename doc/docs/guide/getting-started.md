# 快速下手

::: warning 前提条件
<!--@formatter:off-->
Golang需要1.16及以上版本
<!--@formatter:on-->
:::

本文会帮助你从头搭建一个简单的Golang应用

1. 创建并进入一个新目录

``` bash
mkdir pangu-app && cd pangu-app
```

2. 初始化Go Module

``` bash
go mod init
```

3. 创建启动器

<<< @/../example/bootstrap.go

4. 创建启动文件

<<< @/../example/main.go

5. 启动本运行服务器

``` bash
go build
```
