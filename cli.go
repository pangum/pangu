package pangu

import (
	`github.com/urfave/cli/v2`
)

func newApp() (app *cli.App) {
	app = cli.NewApp()
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	// 定制帮助信息
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "显示所有命令或者帮助信息",
	}

	return
}
