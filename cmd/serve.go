package cmd

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"

	"github.com/pangum/pangu/app"
)

var _ app.Command = (*Serve)(nil)

// Serve 描述一个提供服务的命令
type Serve struct {
	*Command

	beforeExecutors []app.Executor
	afterExecutors  []app.Executor
	serves          []app.Serve

	stop   chan os.Signal
	logger app.Logger
}

// NewServe 创建服务命令
func NewServe(logger app.Logger) *Serve {
	return &Serve{
		Command: New(`serve`, Usage(`启动服务`), Aliases(`s`)),

		beforeExecutors: make([]app.Executor, 0),
		afterExecutors:  make([]app.Executor, 0),
		serves:          make([]app.Serve, 0, 1),

		stop:   make(chan os.Signal),
		logger: logger,
	}
}

func (s *Serve) Adds(components ...interface{}) (err error) {
	for _, component := range components {
		switch typ := component.(type) {
		case app.Serve:
			s.AddServes(typ)
		case app.Executor:
			s.AddExecutors(typ)
		default:
			err = exc.NewField(`不支持的类型`, field.Any(`type`, typ))
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
		s.logger.Info(`启动服务开始`, field.Int(`count`, serveCount))
		if err = s.startServes(ctx); nil != err {
			return
		}
		s.logger.Info(`启动服务成功`, field.Int(`count`, serveCount))
	}

	// 执行生命周期方法
	if 0 != len(s.afterExecutors) {
		if err = app.RunExecutors(s.afterExecutors...); nil != err {
			return
		}
	}

	return
}

func (s *Serve) startServes(ctx *app.Context) (err error) {
	wg := sync.WaitGroup{}
	worker := len(s.serves)
	wg.Add(worker)

	for _, serve := range s.serves {
		_serve := serve
		go func() {
			defer wg.Done()

			s.logger.Info(`启动服务成功`, field.String(`name`, _serve.Name()))
			// 记录时间，如果发生错误的时间小于500毫秒，就是执行错误，应该立即退出；如果大于，则只记录日志
			now := time.Now()
			if startErr := _serve.Start(); nil != startErr {
				errTime := time.Now()
				if errTime.Sub(now) > 500*time.Millisecond {
					s.logger.Info(`服务执行错误`, field.String(`name`, _serve.Name()), field.Error(startErr))
				} else {
					panic(startErr)
				}
			}
		}()
	}

	// 注册退出信号处理
	signal.Notify(s.stop, syscall.SIGTERM, syscall.SIGINT)
	<-s.stop
	if err = s.stopServes(ctx); nil != err {
		panic(err)
	}

	// 等待所有服务启动
	wg.Wait()

	return
}

func (s *Serve) stopServes(_ *app.Context) (err error) {
	wg := sync.WaitGroup{}
	worker := len(s.serves)
	wg.Add(worker)

	for _, serve := range s.serves {
		_serve := serve
		go func() {
			defer wg.Done()

			s.logger.Info(`停止服务成功`, field.String(`name`, _serve.Name()))
			if stopErr := _serve.Stop(); nil != stopErr {
				s.logger.Info(`停止服务出错`, field.String(`name`, _serve.Name()), field.Error(stopErr))
			}
		}()
	}

	// 等待所有程序退出
	wg.Wait()

	return
}
