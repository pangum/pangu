package command

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/runtime"
)

var _ app.Command = (*Serve)(nil)

// Serve 描述一个提供服务的命令
type Serve struct {
	*Default

	serves []app.Serve
	signal chan os.Signal
	logger app.Logger
}

func NewServe(logger app.Logger) *Serve {
	return &Serve{
		Default: New("serve").Usage("启动服务").Aliases("s").Build(),

		serves: make([]app.Serve, 0, 1),
		signal: make(chan os.Signal),
		logger: logger,
	}
}

func (s *Serve) Add(serves ...app.Serve) {
	s.serves = append(s.serves, serves...)
}

func (s *Serve) Run(ctx *runtime.Context) (err error) {
	count := len(s.serves)
	if 0 != count {
		s.logger.Info("启动服务开始", field.New("count", count))
		if err = s.start(ctx); nil != err {
			return
		}
		s.logger.Info("启动服务成功", field.New("count", count))
	}

	return
}

func (s *Serve) start(ctx *runtime.Context) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(s.serves))

	for _, serve := range s.serves {
		go s.startServe(serve, wg, &err)
	}

	// 注册退出信号处理
	signal.Notify(s.signal, syscall.SIGTERM, syscall.SIGINT)
	<-s.signal
	if err = s.stop(ctx); nil != err {
		panic(err)
	}

	// 等待所有服务启动
	wg.Wait()

	return
}

func (s *Serve) startServe(serve app.Serve, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	s.logger.Info("启动服务成功", field.New[string]("name", serve.Name()))
	// 记录时间，如果发生错误的时间小于500毫秒，就是执行错误，应该立即退出；如果大于，则只记录日志
	now := time.Now()
	if se := serve.Start(); nil != se {
		errTime := time.Now()
		if errTime.Sub(now) > 500*time.Millisecond {
			s.logger.Info("服务执行错误", field.New[string]("name", serve.Name()), field.Error(se))
		} else {
			*err = se
		}
	}
}

func (s *Serve) stop(_ *runtime.Context) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(s.serves))

	for _, serve := range s.serves {
		go s.stopServe(serve, wg, &err)
	}

	// 等待所有程序退出
	wg.Wait()

	return
}

func (s *Serve) stopServe(serve app.Serve, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	s.logger.Info("停止服务成功", field.New("name", serve.Name()))
	if se := serve.Stop(); nil != se {
		*err = se
		s.logger.Info("停止服务出错", field.New("name", serve.Name()), field.Error(se))
	}
}
