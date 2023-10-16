package runtime

import (
	"github.com/urfave/cli/v2"
)

type Shell struct {
	*cli.App
}

func NewShell() (shell *Shell) {
	shell = new(Shell)
	shell.App = cli.NewApp()
	shell.App.EnableBashCompletion = true
	shell.App.UseShortOptionHandling = true
	// 定制版本显示，版本号不可改变
	shell.App.Version = Version

	// 定制版本标志
	cli.VersionFlag = &cli.BoolFlag{
		Name: "version",
		Aliases: []string{
			"v",
			"ver",
		},
		Usage: "显示应用程序版本信息",
	}
	// 定制帮助信息
	cli.HelpFlag = &cli.BoolFlag{
		Name: "help",
		Aliases: []string{
			"h",
		},
		Usage: "显示所有命令或者帮助信息",
	}

	return
}
