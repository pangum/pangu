# 应用程序配置

一个应用程序，多半都会有很多配置项，包括并不限于

- Http服务器配置
- Http客户端配置
- 对象存储相关配置
- 腾讯云相关配置
- 阿里云相关配置
- 邮件配置
- 其它配置项

这些配置对于应用程序来说，**很重要又显得不那么重要**

- 重要是因为，应用程序必须有这些配置才能执行
- 不重要是因为这些配置项和业务并不那么相关
- 如果能不关注配置的载入及验证，只需要用，那么一切都将变得美好起来

## 配置

配置是一个简单的`Struct`，这个`Struct`里面定义了应用程序需要的各种配置，建议

- 有一个最外层的配置定义
- 配置定义讲究层次感
- 不要重复定义配置项（能量做到代码复用）

<<< @/../example/conf/config.go

## 数据验证

`盘古`内置了配置的数据验证功能（默认开启，当然也可以关闭数据验证），使用[Golang validator](https://github.com/go-playground/validator)
并在上面做了一些功能扩展，包括

- 手机号验证
- 国际化支持
- 密码
- 文件名
- 其它

这些扩展被集合在了一个叫[Validatorx](https://github.com/storezhang/validatorx)的代码仓库里，可以很方便的扩展验证器
要使用数据验证也很简单，编写对应的`tag`即可，就像这样

<<< @/../example/conf/config.go{16,18}

## 解封配置

意思是将原来的配置内部信息暴露出来，这样调用方就可以只依赖具体的配置而不是将整个配置结构体传过去

<<< @/../example/conf/config.go{29,34}

需要明白的是

- 解封配置并不是必须的
- 如果不解封配置，那么就直接依赖整个配置最外层结构体`Config`
- 解封的深度需要开发者在代码量和使用方便两者中平衡

## 配置文件路径配置

使用`--conf`和`-c`都可以配置配置文件的路径，路径可以是`绝对路径`也可以是`相对路径`，开发者可以视自己的使用情况而定，路径有默认值（即开发者可以不配置路径）
``` shell
./example --conf path/to/your/config/filepath
```

## 默认的配置文件的加载顺序

`盘古`严格按照如下顺序加载配置文件

- 用户通过参数`--conf`或者`-c`指定路径
- ./application.yaml
- ./application.yml
- ./application.toml
- ./application.json
- ./application.xml
- ./conf/application.yaml
- ./conf/application.yml
- ./conf/application.toml
- ./conf/application.json
- ./conf/application.xml
- ./app.yaml
- ./app.yml
- ./app.toml
- ./app.json
- ./app.xml
- ./conf/app.yaml
- ./conf/app.yml
- ./conf/app.toml
- ./conf/app.json
- ./conf/app.xml

如果上面的路径都找不到配置文件（即文件不存在），将会返回路径不存在的错误
