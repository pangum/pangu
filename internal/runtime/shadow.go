package runtime

import (
	"github.com/urfave/cli/v2"
)

type Shadow struct {
	*cli.App
}

func NewShadow() (shadow *Shadow) {
	shadow = new(Shadow)
	shadow.App = cli.NewApp()
	shadow.App.EnableBashCompletion = false
	shadow.App.UseShortOptionHandling = false
	// 对于找不到的命令，暂时不做任何处理
	shadow.App.CommandNotFound = shadow.notfound
	shadow.App.InvalidFlagAccessHandler = shadow.notfound
	shadow.App.HideHelpCommand = true
	shadow.App.HideVersion = true

	return
}

func (s *Shadow) notfound(_ *cli.Context, c string) {}
