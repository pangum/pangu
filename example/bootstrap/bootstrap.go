package bootstrap

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/rest`
)

// 启动器，必须实现pangu.Bootstrap接口
type bootstrap struct {
	application *pangu.Application
}

// NewBootstrap 构造方法，由于Golang语言的限制，只能返回接口类型
// 这里之所以使用公开方法是因为方便代码实现（其它地方要引用启动器，免得每个实例都重新定义一个启动器）
// 正式使用中，建议使用非公开方法来创建启动器
func NewBootstrap(application *pangu.Application) pangu.Bootstrap {
	return &bootstrap{
		application: application,
	}
}

func (b *bootstrap) Setup() (err error) {
	if err = b.application.Invoke(func(server *rest_test.Server) error {
		return b.application.Adds(server)
	}); nil != err {
		return
	}

	return
}
