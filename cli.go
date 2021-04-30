package pangu

import (
	`github.com/urfave/cli/v2`
)

// NewCliApp 创建新的命令行程序
func NewCliApp() (app *cli.App) {
	app = cli.NewApp()
	app.EnableBashCompletion = true

	return
}
