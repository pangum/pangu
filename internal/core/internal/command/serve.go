package command

import (
	"context"
	"sync"
	"time"

	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/boot/internal/internal/command"
	"github.com/harluo/boot/internal/internal/constant"
	"github.com/harluo/boot/internal/internal/param"
)

var _ application.Command = (*Serve)(nil)

// Serve 描述一个提供服务的命令
type Serve struct {
	*command.Default

	serves  []application.Serve
	logger  log.Logger
	exiting bool
}

func NewServe(logger log.Logger) *Serve {
	return &Serve{
		Default: command.NewDefault(&param.Command{
			Name:  constant.CommandServe,
			Usage: `启动服务`,
			Aliases: []string{
				`s`,
				`srv`,
			},
		}),

		serves:  make([]application.Serve, 0, 1),
		logger:  logger,
		exiting: false,
	}
}

func (s *Serve) Add(serves ...application.Serve) {
	s.serves = append(s.serves, serves...)
}

func (s *Serve) Run(ctx context.Context) (err error) {
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

func (s *Serve) start(ctx context.Context, count int) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(count)
	for _, serve := range s.serves {
		cloned := serve
		go s.startServe(ctx, cloned, wg, &err)
	}
	s.logger.Debug("启动所有服务成功", field.New("count", count))
	wg.Wait()

	return
}

func (s *Serve) startServe(ctx context.Context, serve application.Serve, wg *sync.WaitGroup, err *error) {
	defer wg.Done()

	s.logger.Info("启动服务成功", field.New[string]("name", serve.Name()))
	// 记录时间，如果发生错误的时间小于500毫秒，就是执行错误，应该立即退出；如果大于，则只记录日志
	now := time.Now()
	if se := serve.Start(ctx); nil != se && !s.exiting {
		errTime := time.Now()
		if errTime.Sub(now) > 500*time.Millisecond {
			s.logger.Info("服务执行错误", field.New[string]("name", serve.Name()), field.Error(se))
		} else {
			*err = se
		}
	}
}

func (s *Serve) stopServe(ctx context.Context, serve application.Serve, wg *sync.WaitGroup, err *error) {
	defer wg.Done()
	if se := serve.Stop(ctx); nil != se {
		*err = se
		s.logger.Info("停止服务出错", field.New("name", serve.Name()), field.Error(se))
	} else {
		s.logger.Info("停止服务成功", field.New("name", serve.Name()))
	}
}
