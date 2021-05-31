package main

import (
	`github.com/storezhang/pangu`
	`github.com/storezhang/pangu/example/command`
	`github.com/storezhang/pangu/example/conf`
	`github.com/storezhang/pangu/example/rest`
)

// 启动器，必须实现pangu.Bootstrap接口
type bootstrap struct {
	application *pangu.Application
}

// 构造方法，由于Golang语言的限制，只能返回接口类型
func newBootstrap(application *pangu.Application) pangu.Bootstrap {
	return &bootstrap{application: application}
}

func (b *bootstrap) Setup() (err error) {
	if err = b.provides(conf.Provides, rest.Provides, command.Provides); nil != err {
		return
	}
	if err = b.application.Invoke(func(rest *rest.Server) error {
		return b.application.AddServes(rest)
	}); nil != err {
		return
	}

	return
}

func (b *bootstrap) provides(provides ...func(application *pangu.Application) error) (err error) {
	for _, provide := range provides {
		if err = provide(b.application); nil != err {
			break
		}
	}

	return
}
