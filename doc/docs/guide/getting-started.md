# 快速下手

::: warning 前提条件
<!--@formatter:off-->
Golang需要1.16及以上版本
<!--@formatter:on-->
:::

本文会帮助你从头搭建一个简单的Golang应用

## 创建并进入一个新目录

``` shell
mkdir pangu-app && cd pangu-app
```

## 初始化Go Module

``` shell
go mod init
```

## 创建启动器

<<< @/../example/bootstrap.go

## 创建启动文件

<<< @/../example/main.go

## 编译

``` shell
go build
```

## 运行

``` shell
./example serve
```
