# 版本信息

可以很方便的输出如下版本信息
![版本信息](/version.png)

::: warning 注意

有三种方式可以很方便的注入版本信息，推荐在CI/CD系统中自动化注入版本信息
:::

## 在CI/CD系统中注入版本信息

推荐在CI/CD系统中做版本信息的注入，做到自动化

### Drone

``` shell
- name: 编译
image: pangum/pangu
volumes:
  - name: go
    path: /var/lib/go
settings:
  output: pangu

volumes:
  - name: go
    host:
      path: /var/lib/cache/go
```

### Jenkins

``` shell
TODO
```

### TravisCI

``` shell
TODO
```

### Gitlab CI

``` shell
TODO
```

### Github Actions

``` shell
TODO
```

## 编译时注入版本信息

不建议在非CI/CD系统中注入版本信息，因为那太耗时了（费力不讨好）

``` shell
go build -ldflags "-X 'github.com/pangum/pangu.Name=Pangu" -o pangu
```
