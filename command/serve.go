package command

import (
	`sync`

	`github.com/storezhang/glog`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/pangu/app`
)

var _ app.Command = (*Serve)(nil)

type (
	// serve 保持和pangu.Serve接口一致，纯粹是方便外部调用而设
	serve interface {
		// Run 运行服务器
		Run() (err error)
		// Name 服务器名称
		Name() string
	}

	// Serve 描述一个提供服务的命令
	Serve struct {
		Command

		serves      []serve
		serverCount int

		logger glog.Logger
	}
)

// NewServe 创建服务命令
func NewServe(logger glog.Logger) *Serve {
	return &Serve{
		Command: Command{
			Name:    "serve",
			Aliases: []string{"s", "S", "Serve"},
			Usage:   "启动服务",
		},

		// 至少有一个服务器必须加入
		serves:      make([]serve, 0, 1),
		serverCount: 0,

		logger: logger,
	}
}

func (cs *Serve) Add(serve serve) {
	cs.serves = append(cs.serves, serve)
	cs.serverCount++
}

func (cs *Serve) Run(_ *Context) (err error) {
	wg := sync.WaitGroup{}
	worker := cs.serverCount
	wg.Add(worker)

	for _, server := range cs.serves {
		server := server
		go func() {
			defer wg.Done()

			cs.logger.Info("启动服务器成功", field.String("name", server.Name()))
			// 服务器不允许中途有服务器启动错误，如果有，应该立即关掉容器
			// 如果调用者想并行执行，可以使用recover机制来阻止程序退出
			if err = server.Run(); nil != err {
				panic(err)
			}
		}()
	}

	wg.Wait()

	return
}
