package command

import (
	`io/fs`
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

		serves []serve
		// 升级
		migration migration

		logger glog.Logger
	}
)

// NewServe 创建服务命令
func NewServe(logger *glog.ZapLogger) *Serve {
	return &Serve{
		Command: Command{
			Name:    "serve",
			Aliases: []string{"s", "S", "Serve"},
			Usage:   "启动服务",
		},

		serves: make([]serve, 0, 1),
		logger: logger,
	}
}

func (s *Serve) Add(serve serve) {
	s.serves = append(s.serves, serve)
}

func (s *Serve) AddMigration(migration fs.FS) {
	s.migration.addMigration(migration)
}

func (s *Serve) Run(ctx *app.Context) (err error) {
	if s.migration.shouldMigration() {
		s.logger.Info("执行升级开始")

		if err = s.migration.migrate(); nil != err {
			return
		}

		s.logger.Info("执行升级成功")
	}

	serveCount := len(s.serves)
	if 0 != len(s.serves) {
		s.logger.Info("启动服务开始", field.Int("count", serveCount))

		if err = s.runServes(ctx); nil != err {
			return
		}

		s.logger.Info("启动服务成功", field.Int("count", serveCount))
	}

	return
}

func (s *Serve) runServes(_ *app.Context) (err error) {
	wg := sync.WaitGroup{}
	worker := len(s.serves)
	wg.Add(worker)

	for _, serve := range s.serves {
		serve := serve
		go func() {
			defer wg.Done()

			s.logger.Info("启动服务器成功", field.String("name", serve.Name()))
			// 服务器不允许中途有服务器启动错误，如果有，应该立即关掉容器
			// 如果调用者想并行执行，可以使用recover机制来阻止程序退出
			if err = serve.Run(); nil != err {
				panic(err)
			}
		}()
	}

	wg.Wait()

	return
}
