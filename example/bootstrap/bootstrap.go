package bootstrap

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/rest`
)

// 启动器，必须实现pangu.Bootstrap接口
type (
	bootstrap struct {
		application *pangu.Application
		rest        *rest.Server
	}

	bootstrapIn struct {
		pangu.In

		Application *pangu.Application
		Rest        *rest.Server
	}
)

// NewBootstrap 构造方法，由于Golang语言的限制，只能返回接口类型
// 这里之所以使用公开方法是因为方便代码实现（其它地方要引用启动器，免得每个实例都重新定义一个启动器）
// 正式使用中，建议使用非公开方法来创建启动器
// bootstrapIn不是必须的，这里是为了演示过多依赖项才这么做，一般来说，超过4个依赖建议使用 pangu.In，不然方法参数太多代码不好看
func NewBootstrap(in bootstrapIn) pangu.Bootstrap {
	return &bootstrap{
		application: in.Application,
		rest:        in.Rest,
	}
}

func (b *bootstrap) Setup() (err error) {
	return b.application.AddServes(b.rest)
}
