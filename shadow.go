package pangu

import (
	"github.com/urfave/cli/v2"
)

func newShadow() (shadow *cli.App) {
	shadow = cli.NewApp()
	shadow.EnableBashCompletion = false
	shadow.UseShortOptionHandling = false
	// 对于找不到的命令，暂时不做任何处理
	shadow.CommandNotFound = func(ctx *cli.Context, command string) {}

	return
}
