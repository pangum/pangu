package runtime

import (
	"github.com/pangum/pangu/internal"
	"github.com/urfave/cli/v2"
)

type Shell struct {
	*cli.App
}

func NewShell() (shell *Shell) {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	// 定制版本显示，版本号不可改变
	app.Version = internal.Version
	// 对于找不到的命令，暂时不做任何处理
	app.CommandNotFound = shell.notfound

	shell = new(Shell)
	shell.App = app

	return
}

func (s *Shell) notfound(_ *cli.Context, _ string) {}
