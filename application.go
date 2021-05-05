package pangu

import (
	`io/fs`

	`github.com/storezhang/glog`
	`github.com/storezhang/pangu/app`
	`github.com/storezhang/pangu/command`
	`github.com/urfave/cli/v2`
)

// Application 应用程序，可以加入两种种类型的程序
// Serve 用于描述应用程序内的服务
// Command 用于描述应用程序内可以被执行的命令
type Application struct {
	commands []app.Command
	args     []app.Arg

	// 服务执行器
	serve *command.Serve

	cli *cli.App

	logger glog.Logger
}

// NewApplication 创建一个应用程序
func NewApplication(logger glog.Logger, cli *cli.App) (application Application) {
	application = Application{
		logger: logger,
		cli:    cli,
	}

	return
}

// AddServe 添加一个服务器到应用程序中
func (a *Application) AddServe(serve Serve) {
	a.serve.Add(serve)
}

// AddCommand 添加一个可以被执行的命令到应用程序中
func (a *Application) AddCommand(command app.Command) {
	a.commands = append(a.commands, command)
}

// AddMigration 添加一个服务器到应用程序中
func (a *Application) AddMigration(migration fs.FS) {
	a.serve.AddMigration(migration)
}

// Setup 启动应用程序，提供服务
func (a *Application) Setup() {
	a.addCommands()
	a.addArgs()
}

func (a *Application) addCommands() {
	for _, c := range a.commands {
		// 如果不这样赋值，会出现只执行最后一个命令的Bug，这是Golang语法上的问题
		cmd := c
		a.cli.Commands = append(a.cli.Commands, &cli.Command{
			Name:    cmd.GetName(),
			Aliases: cmd.GetAliases(),
			Usage:   cmd.GetUsage(),
			Action: func(ctx *cli.Context) error {
				return cmd.Run(app.NewContext(ctx))
			},
		})
	}
}

func (a *Application) addArgs() {
	for _, argument := range a.args {
		// 如果不这样赋值，会出现只执行最后一个命令的Bug，这是Golang语法上的问题
		parameter := argument
		switch parameter.GetValue().(type) {
		case string:
			a.cli.Flags = append(a.cli.Flags, &cli.StringFlag{
				Name:        parameter.GetName(),
				Aliases:     parameter.GetAliases(),
				Usage:       parameter.GetUsage(),
				Value:       parameter.GetValue().(string),
				DefaultText: parameter.GetDefaultText(),
			})
		case bool:
			a.cli.Flags = append(a.cli.Flags, &cli.BoolFlag{
				Name:        parameter.GetName(),
				Aliases:     parameter.GetAliases(),
				Usage:       parameter.GetUsage(),
				Value:       parameter.GetValue().(bool),
				DefaultText: parameter.GetDefaultText(),
			})
		case int:
			a.cli.Flags = append(a.cli.Flags, &cli.IntFlag{
				Name:        parameter.GetName(),
				Aliases:     parameter.GetAliases(),
				Usage:       parameter.GetUsage(),
				Value:       parameter.GetValue().(int),
				DefaultText: parameter.GetDefaultText(),
			})
		}
	}
}
