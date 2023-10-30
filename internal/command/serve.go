package command

import (
	"context"
	"sync"
	"time"

	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/runtime"
)

var _ app.Command = (*Serve)(nil)

// Serve 描述一个提供服务的命令
type Serve struct {
	*Default

	serves  []app.Serve
	logger  log.Logger
	exiting bool
}

func NewServe(logger log.Logger) *Serve {
	return &Serve{
		Default: New("serve").Usage("启动服务").Aliases("s").Build(),

		serves:  make([]app.Serve, 0, 1),
		logger:  logger,
		exiting: false,
	}
}

func (s *Serve) Add(serves ...app.Serve) {
	s.serves = append(s.serves, serves...)
}

func (s *Serve) Run(ctx *runtime.Context) (err error) {
	count := len(s.serves)
	if 0 != count {
		s.logger.Debug("启动所有服务开始", field.New("count", count))
		err = s.start(ctx, count)
	}

	return
}

func (s *Serve) Stop(ctx context.Context) (err error) {
	s.exiting = true
	wg := new(sync.WaitGroup)
	wg.Add(len(s.serves))
	for _, serve := range s.serves {
		go s.stopServe(ctx, serve, wg, &err)
	}
	wg.Wait()

	return
}

func (s *Serve) Before(ctx context.Context) (err error) {
	for _, serve := range s.serves {
		err = serve.Before(ctx)
		if nil != err {
			break
		}
	}

	return
}

func (s *Serve) After(ctx context.Context) (err error) {
	for _, serve := range s.serves {
		err = serve.After(ctx)
		if nil != err {
			break
		}
	}

	return
}

func (s *Serve) start(_ *runtime.Context, count int) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(count)
	for _, serve := range s.serves {
		cloned := serve
		go s.startServe(cloned, wg, &err)
	}
	s.logger.Debug("启动所有服务成功", field.New("count", count))
	wg.Wait()

	return
}

func (s *Serve) startServe(serve app.Serve, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	s.logger.Info("启动服务成功", field.New[string]("name", serve.Name()))
	// 记录时间，如果发生错误的时间小于500毫秒，就是执行错误，应该立即退出；如果大于，则只记录日志
	now := time.Now()
	if se := serve.Start(); nil != se && !s.exiting {
		errTime := time.Now()
		if errTime.Sub(now) > 500*time.Millisecond {
			s.logger.Info("服务执行错误", field.New[string]("name", serve.Name()), field.Error(se))
		} else {
			*err = se
		}
	}
}

func (s *Serve) stopServe(ctx context.Context, serve app.Serve, wg *sync.WaitGroup, err *error) {
	defer wg.Done()
	if se := serve.Stop(ctx); nil != se {
		*err = se
		s.logger.Info("停止服务出错", field.New("name", serve.Name()), field.Error(se))
	} else {
		s.logger.Info("停止服务成功", field.New("name", serve.Name()))
	}
}
