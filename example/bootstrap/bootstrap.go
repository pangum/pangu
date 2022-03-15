package bootstrap

import (
	`embed`
	`io/fs`

	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example`
	`github.com/pangum/pangu/example/command`
	`github.com/pangum/pangu/example/conf`
	`github.com/pangum/pangu/example/rest`
)

//go:embed ../db/migration
var migration embed.FS

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
	if err = b.provides(conf.Provides, rest.Provides, command.Provides); nil != err {
		return
	}

	// go embed会带上目录的所有层级，必须将最外层目录全部剥离
	var migrations fs.FS
	if migrations, err = fs.Sub(migration, "db/migration"); nil != err {
		return
	}
	if err = b.application.AddMigration(migrations); nil != err {
		return
	}

	if err = b.application.Invoke(func(in example.componentIn) error {
		return b.application.Adds(in.Rest, in.Test)
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
