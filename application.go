package pangu

import (
	`encoding/json`
	`io/fs`
	`io/ioutil`
	`os`
	`path/filepath`
	`strings`

	`github.com/pelletier/go-toml`
	`github.com/pkg/errors`
	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
	`github.com/storezhang/pangu/app`
	`github.com/storezhang/pangu/arg`
	`github.com/storezhang/pangu/command`
	`github.com/urfave/cli/v2`
	`gopkg.in/yaml.v3`
	`xorm.io/xorm`
)

// Application 应用程序，可以加入两种种类型的程序
// Serve 用于描述应用程序内的服务
// Command 用于描述应用程序内可以被执行的命令
type Application struct {
	commands []app.Command
	args     []app.Arg

	// 存储配置
	config interface{}

	// 内置命令
	serve   *command.Serve
	version *command.Version

	// 应用程序本身的配置
	// 徽标
	banner string

	logger    glog.Logger
	migration migration
}

// NewApplication 创建一个应用程序
func NewApplication(
	logger glog.Logger,
	database gox.DatabaseConfig,
	serve *command.Serve, version *command.Version,
	engine *xorm.Engine,
) (application Application) {
	application = Application{
		serve:   serve,
		version: version,

		logger: logger,
		migration: migration{
			database: database,
			engine:   engine,
		},
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

// SetConfig 添加一个服务器到应用程序中
func (a *Application) SetConfig(config interface{}) {
	a.config = config
}

// GetConfig 取得解析后的配置
func (a *Application) GetConfig() interface{} {
	return a.config
}

// SetBanner 设置徽标
func (a *Application) SetBanner(banner string) {
	a.banner = banner
}

// Run 启动应用程序，提供服务
func (a *Application) Run() (err error) {
	application := cli.NewApp()
	application.EnableBashCompletion = true
	application.Commands = a.parseCommands()
	application.Flags = a.parseArgs()
	application.Action = a.loadConfig

	/*if a.migration.shouldMigration() {
		a.logger.Info("执行升级开始")

		if err = a.migration.migrate(); nil != err {
			return
		}

		a.logger.Info("执行升级成功")
	}*/

	err = application.Run(os.Args)

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

func (a *Application) loadConfig(ctx *cli.Context) (err error) {
	var conf string
	if conf, err = a.findConfigFilepath(ctx.String("conf")); nil != err {
		return
	}

	var data []byte
	if data, err = ioutil.ReadFile(conf); nil != err {
		return
	}

	switch strings.ToLower(filepath.Ext(conf)) {
	case "yml":
	case "yaml":
		err = yaml.Unmarshal(data, &a.config)
	case "json":
		err = json.Unmarshal(data, &a.config)
	case "toml":
		err = toml.Unmarshal(data, &a.config)
	default:
		err = yaml.Unmarshal(data, &a.config)
	}

	return
}

func (a *Application) findConfigFilepath(conf string) (filepath string, err error) {
	filepath = conf
	if "" == filepath {
		filepath = "./application.yml"
	}
	if gox.IsFileExist(filepath) {
		return
	}

	filepath = "./application.yaml"
	if gox.IsFileExist(filepath) {
		return
	}

	filepath = "./application.toml"
	if gox.IsFileExist(filepath) {
		return
	}

	filepath = "./conf/application.yml"
	if gox.IsFileExist(filepath) {
		return
	}

	filepath = "./conf/application.yaml"
	if gox.IsFileExist(filepath) {
		return
	}
	err = errors.New("找不到配置文件")

	return
}
