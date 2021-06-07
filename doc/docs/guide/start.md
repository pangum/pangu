# 快速下手

::: warning 前提条件
<!--@formatter:off-->
Golang需要1.16及以上版本
<!--@formatter:on-->
:::

本文会帮助你从头搭建一个简单的Golang应用

## 新建项目

创建普通的Golang项目，容纳项目代码

``` shell
mkdir pangu-app && cd pangu-app
```

## 初始化`Go Module`

`盘古`建议使用`Go Module`安装，其它方式不在支持之列

``` shell
go mod init
```

## 创建启动器

启动器是`盘古`最重要的部件，一个好的启动器能让项目代码结构变得简单

<<< @/../example/bootstrap.go

## 创建启动文件

`盘古`推行瘦启动器（和`SpringBoot`一样），只有少量的代码（大部分是配置）

<<< @/../example/main.go

## 编译

编译Golang代码为可执行程序，因为是运行时依赖注入，所以不需要增加额外依赖代码

``` shell
go build
```

## 运行

运行可执行程序，你就能看到程序的输出啦

``` shell
./example serve
```
