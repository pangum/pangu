package command

import (
	`errors`
	`sync`

	`github.com/storezhang/gox/field`
	`github.com/storezhang/pangu/app`
)

var _ app.Command = (*Serve)(nil)

// Serve 描述一个提供服务的命令
type Serve struct {
	Base

	beforeExecutors []app.Executor
	afterExecutors  []app.Executor
	serves          []app.Serve
	logger          app.Logger
}

// NewServe 创建服务命令
func NewServe(logger app.Logger) *Serve {
	return &Serve{
		Base: Base{
			name:    "serve",
			aliases: []string{"s"},
			usage:   "启动服务",
		},

		beforeExecutors: make([]app.Executor, 0, 0),
		afterExecutors:  make([]app.Executor, 0, 0),
		serves:          make([]app.Serve, 0, 1),
		logger:          logger,
	}
}

func (s *Serve) Adds(components ...interface{}) (err error) {
	for _, component := range components {
		switch component.(type) {
		case app.Serve:
			s.AddServes(component.(app.Serve))
		case app.Executor:
			s.AddExecutors(component.(app.Executor))
		default:
			err = errors.New("不支持的类型")
		}

		if nil != err {
			break
		}
	}

	return
}

func (s *Serve) AddServes(serves ...app.Serve) {
	s.serves = append(s.serves, serves...)
}

func (s *Serve) AddExecutors(executors ...app.Executor) {
	for _, executor := range executors {
		switch executor.Type() {
		case app.ExecutorTypeBeforeServe:
			s.beforeExecutors = append(s.beforeExecutors, executor)
		case app.ExecutorTypeAfterServe:
			s.afterExecutors = append(s.afterExecutors, executor)
		}
	}
}

func (s *Serve) Run(ctx *app.Context) (err error) {
	// 执行生命周期方法
	if 0 != len(s.beforeExecutors) {
		if err = app.RunExecutors(s.beforeExecutors...); nil != err {
			return
		}
	}

	serveCount := len(s.serves)
	if 0 != serveCount {
		s.logger.Info("启动服务开始", field.Int("count", serveCount))
		if err = s.runServes(ctx); nil != err {
			return
		}
		s.logger.Info("启动服务成功", field.Int("count", serveCount))
	}

	// 执行生命周期方法
	if 0 != len(s.afterExecutors) {
		if err = app.RunExecutors(s.afterExecutors...); nil != err {
			return
		}
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
