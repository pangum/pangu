package runtime

import (
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/pangum/pangu/internal"
	"github.com/urfave/cli/v2"
)

type Shell struct {
	*cli.App

	logger log.Logger
}

func NewShell() (shell *Shell) {
	shell = new(Shell)
	shell.App = cli.NewApp()
	shell.App.EnableBashCompletion = true
	shell.App.UseShortOptionHandling = true
	// 定制版本显示，版本号不可改变
	shell.App.Version = internal.Version
	// 对于找不到的命令，暂时不做任何处理
	shell.App.CommandNotFound = shell.notfound

	return
}

func (s *Shell) Logger(logger log.Logger) {
	s.logger = logger
}

func (s *Shell) notfound(_ *cli.Context, command string) {
	s.logger.Warn("你调用了不存在的命令", field.New("command", command))
}
