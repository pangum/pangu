# 版本信息

可以很方便的输出如下版本信息
![version](/version.png)

## 在CI/CD中注入版本信息（推荐）

推荐在CI/CD系统中做版本信息的注入，做到自动化

### Drone

``` shell
- name: 编译
image: golang
pull: if-not-exists
volumes:
  - name: deps
    path: /go
commands:
  - export ldflags="-s"
  - # 注入版本信息
  - ldflags="$ldflags -X 'github.com/storezhang/pangu.AppName=Archtech服务器'"
  - ldflags="$ldflags -X 'github.com/storezhang/pangu.AppVersion=${DRONE_TAG=$DRONE_COMMIT_BRANCH:latest}'"
  - ldflags="$ldflags -X 'github.com/storezhang/pangu.BuildVersion=$$DRONE_BUILD_NUMBER'"
  - ldflags="$ldflags -X 'github.com/storezhang/pangu.BuildTime=$(TZ="Asia/Shanghai" date --date "@$$DRONE_BUILD_STARTED" +"%F %T")'"
  - ldflags="$ldflags -X 'github.com/storezhang/pangu.ScmRevision=$$DRONE_COMMIT_SHA'"
  - ldflags="$ldflags -X 'github.com/storezhang/pangu.ScmBranch=$$DRONE_COMMIT_BRANCH'"
  - # 编译
  - CGO_ENABLED=0 GOOS=linux go build -ldflags "$ldflags" -o archtech
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
go build -ldflags "-X 'github.com/storezhang/pangu.AppName=Archtech服务器" -o archtech
```
