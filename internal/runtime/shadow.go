package runtime

import (
	"github.com/urfave/cli/v2"
)

type Shadow struct {
	*cli.App
}

func NewShadow() (shadow *Shadow) {
	shadow = new(Shadow)

	app := cli.NewApp()
	app.EnableBashCompletion = false
	app.UseShortOptionHandling = false
	// 对于找不到的命令，暂时不做任何处理
	app.CommandNotFound = shadow.notfound

	shadow.App = app

	return
}

func (s *Shadow) notfound(_ *cli.Context, _ string) {}
