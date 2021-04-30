package pangu

import (
	`io/fs`
	`os`

	`github.com/storezhang/glog`
	`github.com/storezhang/pangu/app`
	`github.com/storezhang/pangu/arg`
	`github.com/storezhang/pangu/command`
	`github.com/urfave/cli/v2`
)

// Application 应用程序，可以加入两种种类型的程序
// Serve 用于描述应用程序内的服务
// Command 用于描述应用程序内可以被执行的命令
type Application struct {
	commands []app.Command
	args     []app.Arg

	// 内置命令
	serve   *command.Serve
	version *command.Version

	// 应用程序本身的配置
	// 徽标
	banner string

	cli *cli.App

	logger    glog.Logger
	migration migration
}

// NewApplication 创建一个应用程序
func NewApplication(
	logger glog.Logger,
	cli *cli.App,
// database gox.DatabaseConfig,
	serve *command.Serve, version *command.Version,
// engine *xorm.Engine,
) (application Application) {
	application = Application{
		serve:   serve,
		version: version,

		logger: logger,
		cli:    cli,
		/*migration: migration{
			database: database,
			engine:   engine,
		},*/
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
	a.migration.addMigration(migration)
}

// SetBanner 设置徽标
func (a *Application) SetBanner(banner string) {
	a.banner = banner
}

// Run 启动应用程序，提供服务
func (a *Application) Run() (err error) {
	a.cli.Commands = a.parseCommands()
	a.cli.Flags = a.parseArgs()

	/*if a.migration.shouldMigration() {
		a.logger.Info("执行升级开始")

		if err = a.migration.migrate(); nil != err {
			return
		}

		a.logger.Info("执行升级成功")
	}*/

	err = a.cli.Run(os.Args)

	return
}

func (a *Application) parseCommands() (commands []*cli.Command) {
	// 添加内置命令
	a.commands = append(a.commands, a.serve, a.version)

	commands = make([]*cli.Command, 0, len(a.commands))
	for _, c := range a.commands {
		// 如果不这样赋值，会出现只执行最后一个命令的Bug，这是Golang语法上的问题
		cmd := c
		cc := &cli.Command{
			Name:    cmd.GetName(),
			Aliases: cmd.GetAliases(),
			Usage:   cmd.GetUsage(),
			Action: func(ctx *cli.Context) error {
				return cmd.Run(app.NewContext(ctx))
			},
		}
		commands = append(commands, cc)
	}

	return
}

func (a *Application) parseArgs() (flags []cli.Flag) {
	// 添加内置参数
	a.args = append(a.args, &arg.String{
		Arg: arg.Arg{
			Name:        "conf",
			Aliases:     []string{"c", "config", "configuration"},
			Usage:       "",
			DefaultText: "./conf/application.yaml",
		},
		Value: "./conf/application.yaml",
	})

	flags = make([]cli.Flag, 0, len(a.args))
	for _, argument := range a.args {
		// 如果不这样赋值，会出现只执行最后一个命令的Bug，这是Golang语法上的问题
		parameter := argument
		switch parameter.GetValue().(type) {
		case string:
			flags = append(flags, &cli.StringFlag{
				Name:        parameter.GetName(),
				Aliases:     parameter.GetAliases(),
				Usage:       parameter.GetUsage(),
				Value:       parameter.GetValue().(string),
				DefaultText: parameter.GetDefaultText(),
			})
		case bool:
			flags = append(flags, &cli.BoolFlag{
				Name:        parameter.GetName(),
				Aliases:     parameter.GetAliases(),
				Usage:       parameter.GetUsage(),
				Value:       parameter.GetValue().(bool),
				DefaultText: parameter.GetDefaultText(),
			})
		}
	}

	return
}
