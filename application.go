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

	`github.com/mcuadros/go-defaults`
	`github.com/pelletier/go-toml`
	`github.com/storezhang/glog`
	`github.com/storezhang/gox`
	`github.com/storezhang/pangu/app`
	`github.com/storezhang/pangu/command`
	`github.com/storezhang/validatorx`
	`github.com/urfave/cli/v2`
	`go.uber.org/dig`
	`gopkg.in/yaml.v3`
)

// Application 应用程序，可以加入两种种类型的程序
// Serve 用于描述应用程序内的服务
// Command 用于描述应用程序内可以被执行的命令
type Application struct {
	options options

	container *dig.Container
}

var (
	once        sync.Once
	application *Application
)

// New 创建一个应用程序
// 使用单例模式
func New(opts ...option) *Application {
	once.Do(func() {
		application = &Application{
			options: defaultOptions(),

			container: dig.New(),
		}
	})

	for _, opt := range opts {
		opt.apply(&application.options)
	}

	return application
}

// AddServes 添加一个服务器到应用程序中
func (a *Application) AddServes(serves ...Serve) error {
	return a.container.Invoke(func(cmd *command.Serve) {
		for _, serve := range serves {
			// 为了防止包循环引用不得已的办法
			cmd.Adds(serve)
		}
	})
}

// AddCommands 添加一个可以被执行的命令到应用程序中
func (a *Application) AddCommands(commands ...app.Command) error {
	return a.container.Invoke(func(startup *cli.App) {
		for _, cmd := range commands {
			cmd := cmd
			startup.Commands = append(startup.Commands, &cli.Command{
				Name:    cmd.Name(),
				Aliases: cmd.Aliases(),
				Usage:   cmd.Usage(),
				Action: func(ctx *cli.Context) error {
					return cmd.Run(app.NewContext(ctx))
				},
			})
		}
	})
}

// AddArgs 添加参数
func (a *Application) AddArgs(args ...app.Arg) error {
	return a.container.Invoke(func(startup *cli.App) {
		for _, argument := range args {
			parameter := argument
			startup.Flags = append(startup.Flags, parameter.ParseFlag())
		}
	})
}

// AddMigration 添加一个升级脚本到系统中
func (a *Application) AddMigration(source fs.FS) error {
	return a.container.Invoke(func(migration *migration) {
		migration.addSource(source)
	})
}

func (a *Application) Set(constructor interface{}) (err error) {
	return a.container.Provide(constructor)
}

func (a *Application) Sets(constructors ...interface{}) (err error) {
	for _, constructor := range constructors {
		if err = a.container.Provide(constructor); nil != err {
			return
		}
	}

	return
}

func (a *Application) Get(function interface{}) error {
	return a.container.Invoke(function)
}

// Run 启动应用程序
func (a *Application) Run(bootstrap func(*Application) Bootstrap) (err error) {
	// 输出标志信息
	if "" != a.options.banner.data {
		if err = a.options.banner.print(); nil != err {
			return
		}
	}

	if err = a.addProvides(); nil != err {
		return
	}
	if err = a.setup(); nil != err {
		return
	}

	// 添加启动器到依赖关系中
	if err = a.Set(bootstrap); nil != err {
		return
	}

	// 增加内置的命令及参数
	if err = a.addInternalFlags(); nil != err {
		return
	}
	if err = a.addInternalCommands(); nil != err {
		return
	}

	// 加载用户启动器并做好配置
	if err = a.Get(func(bootstrap Bootstrap) error {
		return bootstrap.Setup()
	}); nil != err {
		return
	}

	// 启动应用
	err = a.Get(func(startup *cli.App) error {
		return startup.Run(os.Args)
	})

	// 退出程序，解决最外层panic报错的问题
	// 原理：如果到这个地方还没有发生错误，程序正常退出，外层panic得不到执行
	// 如果发生错误，则所有代码都会返回error直到panic检测到，然后程序整体panic
	os.Exit(0)

	return
}

// GetConfig 取得解析后的配置
func (a *Application) GetConfig(config interface{}) (err error) {
	var (
		once sync.Once
		path *string
	)

	// 参数不允许重复定义，只能执行一次
	once.Do(func() {
		path = flag.String("conf", "./conf/application.yaml", "指定配置文件路径")
		flag.StringVar(path, "c", *path, "指定配置文件路径")
		flag.Parse()
	})

	if reflect.ValueOf(config).Kind() == reflect.Ptr {
		err = a.loadConfig(config, *path)
	} else {
		err = a.loadConfig(&config, *path)
	}

	return
}

func (a *Application) setup() error {
	cli.AppHelpTemplate = a.options.helpAppTemplate
	cli.CommandHelpTemplate = a.options.helpCommandTemplate
	cli.SubcommandHelpTemplate = a.options.helpSubcommandTemplate
	cli.VersionPrinter = func(ctx *cli.Context) {
		_ = a.Get(func(version *command.Version) error {
			return version.Run(app.NewContext(ctx))
		})
	}

	return a.Get(func(startup *cli.App) {
		startup.Name = a.options.name
		startup.Description = a.options.description
		startup.Usage = a.options.usage
		startup.Copyright = a.options.copyright
		if 0 != len(a.options.authors) {
			authors := make([]*cli.Author, 0, len(a.options.authors))
			for _, author := range a.options.authors {
				authors = append(authors, &cli.Author{Name: author.Name, Email: author.Email})
			}
			startup.Authors = authors
		}
	})
}

func (a *Application) addInternalCommands() error {
	type commandsIn struct {
		In

		Serve     *command.Serve
		Version   *command.Version
		Migrate   *command.Migrate
		Migration *migration
	}

	return a.Get(func(in commandsIn) error {
		in.Serve.SetMigration(in.Migration)
		in.Migrate.SetMigration(in.Migration)

		return a.AddCommands(in.Serve, in.Migrate, in.Version)
	})
}

func (a *Application) addInternalFlags() error {
	return a.Get(func(startup *cli.App) {
		startup.Flags = append(startup.Flags, &cli.StringFlag{
			Name:        "conf",
			Aliases:     []string{"c"},
			Usage:       "指定配置文件路径",
			Value:       "./conf/application.yaml",
			DefaultText: "./conf/application.yaml",
		})
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

	// 处理默认值
	if a.options.isDefault {
		defaults.SetDefaults(config)
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
	if nil != err {
		return
	}

	// 验证数据
	if a.options.isValidate {
		err = validatorx.Struct(config)
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

func (a *Application) addProvides() (err error) {
	if err = a.Sets(glog.NewLogger, gox.NewSnowflake, NewResty); nil != err {
		return
	}
	if err = a.Sets(command.NewServe, command.NewVersion, command.NewMigrate); nil != err {
		return
	}
	if err = a.Sets(appName, appVersion, buildVersion, buildTime, scmRevision, scmBranch, goVersion); nil != err {
		return
	}
	if err = a.Sets(newApp, newMigration, newZapLogger, app.NewDefaultService); nil != err {
		return
	}

	// 注入快捷方式
	// 解包Http对象
	if err = a.Sets(getClientConfig, getServerConfig); nil != err {
		return
	}

	// 注入自身
	if err = a.Set(func() *Application {
		return a
	}); nil != err {
		return
	}

	return
}
