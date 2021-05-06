package pangu

import (
	`encoding/json`
	`errors`
	`flag`
	`io/fs`
	`io/ioutil`
	`os`
	`path/filepath`
	`reflect`
	`strings`
	`sync`

	`github.com/pelletier/go-toml`
	`github.com/storezhang/gox`
	`github.com/storezhang/pangu/app`
	`github.com/storezhang/pangu/command`
	`github.com/urfave/cli/v2`
	`gopkg.in/yaml.v3`
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
}

// NewApplication 创建一个应用程序
func NewApplication(serve *command.Serve, version *command.Version, cli *cli.App) (application Application) {
	application = Application{
		serve:   serve,
		version: version,

		cli: cli,
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

// GetConfig 取得解析后的配置
func (a *Application) GetConfig(config interface{}) (err error) {
	var (
		once sync.Once
		path *string
	)

	once.Do(func() {
		path = flag.String("conf", "./conf/application.yaml", "指定配置文件路径")
		flag.StringVar(path, "c", *path, "指定配置文件路径")
		flag.StringVar(path, "config", *path, "指定配置文件路径")
		flag.StringVar(path, "configuration", *path, "指定配置文件路径")
		flag.Parse()
	})

	if reflect.ValueOf(config).Kind() == reflect.Ptr {
		err = a.loadConfig(config, *path)
	} else {
		err = a.loadConfig(&config, *path)
	}

	return
}

// Run 启动应用程序，提供服务
func (a *Application) Run() error {
	// 添加内置命令
	a.addInternalFlags()
	a.addInternalCommands()
	// 添加其它命令及参数
	a.addCommands()
	a.addArgs()

	return a.cli.Run(os.Args)
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

func (a *Application) addInternalCommands() {
	a.cli.Commands = append(a.cli.Commands, &cli.Command{
		Name:    a.serve.GetName(),
		Aliases: a.serve.GetAliases(),
		Usage:   a.serve.GetUsage(),
		Action: func(ctx *cli.Context) error {
			return a.serve.Run(app.NewContext(ctx))
		},
	})
	a.cli.Commands = append(a.cli.Commands, &cli.Command{
		Name:    a.version.GetName(),
		Aliases: a.version.GetAliases(),
		Usage:   a.version.GetUsage(),
		Action: func(ctx *cli.Context) error {
			return a.version.Run(app.NewContext(ctx))
		},
	})
}

func (a *Application) addInternalFlags() {
	a.cli.Flags = append(a.cli.Flags, &cli.StringFlag{
		Name:        "conf",
		Aliases:     []string{"c", "config", "configuration"},
		Usage:       "",
		Value:       "./conf/application.yaml",
		DefaultText: "./conf/application.yaml",
	})
}

func (a *Application) loadConfig(config interface{}, path string) (err error) {
	if path, err = a.findConfigFilepath(path); nil != err {
		return
	}

	var data []byte
	if data, err = ioutil.ReadFile(path); nil != err {
		return
	}

	switch strings.ToLower(filepath.Ext(path)) {
	case "yml":
		fallthrough
	case "yaml":
		err = yaml.Unmarshal(data, config)
	case "json":
		err = json.Unmarshal(data, config)
	case "toml":
		err = toml.Unmarshal(data, config)
	default:
		err = yaml.Unmarshal(data, config)
	}

	return
}

func (a *Application) findConfigFilepath(conf string) (path string, err error) {
	path = conf
	if "" == path {
		path = "./application.yml"
	}
	if gox.IsFileExist(path) {
		return
	}

	path = "./application.yaml"
	if gox.IsFileExist(path) {
		return
	}

	path = "./application.toml"
	if gox.IsFileExist(path) {
		return
	}

	path = "./conf/application.yml"
	if gox.IsFileExist(path) {
		return
	}

	path = "./conf/application.yaml"
	if gox.IsFileExist(path) {
		return
	}
	err = errors.New("找不到配置文件")

	return
}
