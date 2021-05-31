# 配置选项

## 配置名称

可以配置应用程序显示帮助信息的名称

``` go
pangu.New(pangu.Name("example"))
```

## 配置徽标

徽标用于应用程序启动时的显示，给用户一种`高大上`的感觉，盘古支持以下几种类型的徽标

- 文本文件
- 图片文件
- 字符串
- 自动生成
- 二进制数据（对于Golang 1.16的静态资源嵌入功能，转换成[]byte）
- 文件数据（对于Golang 1.16的静态资源嵌入功能，转换成embed.fs）

### 文本文件

``` go
pangu.New(pangu.Banner("example", pangu.BannerTypeTxt))
```

### 图片文件

``` go
pangu.New(pangu.Banner("example", pangu.BannerTypeFilepath))
```

### 字符串

``` go
pangu.New(pangu.Banner("example", pangu.BannerTypeString))
```

### 自动生成

``` go
pangu.New(pangu.Banner("example", pangu.BannerTypeAscii))
```

### 二进制数据

``` go
pangu.New(pangu.Banner("example", pangu.BannerTypeBinary))
```

### 文件数据

``` go
pangu.New(pangu.Banner("example", pangu.BannerTypeFile))
```
