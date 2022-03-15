# 徽标

徽标用于应用程序启动时的显示，给用户一种`高大上`的感觉，盘古支持以下几种类型的徽标

- 文本文件
- 图片文件
- 字符串
- 自动生成
- 二进制数据（对于Golang 1.16的静态资源嵌入功能，转换成[]byte）
- 文件数据（对于Golang 1.16的静态资源嵌入功能，转换成embed.fs）

## 文本文件

<<< @/../example/banner/txt.go

## 图片文件

<<< @/../example/banner/filepath.go

## 字符串

<<< @/../example/banner/string.go

## 二进制数据

<<< @/../example/banner/binary.go

## 文件数据

<<< @/../example/banner/file.go
