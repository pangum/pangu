package pangu

import (
	`errors`
	`os`
	`sync`

	`github.com/storezhang/gox`
	`github.com/storezhang/pangu/app`
	`github.com/storezhang/pangu/command`
	`github.com/urfave/cli/v2`
	`go.uber.org/dig`
)

// Application 应用程序，可以加入两种种类型的程序
// Serve 用于描述应用程序内的服务
// Command 用于描述应用程序内可以被执行的命令
type Application struct {
	config    *Config
	options   *options
	container *dig.Container

	beforeExecutors []app.Executor
	afterExecutors  []app.Executor
}

var (
	application *Application
	once        sync.Once
)

// New 创建一个应用程序
// 使用单例模式
func New(opts ...option) *Application {
	once.Do(func() {
		application = &Application{
			options:   defaultOptions(),
			container: dig.New(),

			beforeExecutors: make([]app.Executor, 0, 0),
			afterExecutors:  make([]app.Executor, 0, 0),
		}
		// 注入配置对象，后续使用
		application.config = &Config{
			application: application,
		}

		// 初始化内置变量及内置命令
		// 之所以在这儿完全初始化，是因为可能会在任意地方调用方法，而在调用这些方法时要求必须完成初始化
		// 从另外一个角度上来说，如果初始化失败，panic是被允许的（盘古是应用程序的基础，基础初始化有错误，不panic更待何时）
		// Adds
		// AddCommands
		// AddExecutors
		// AddArgs
		// AddFlags
		if err := application.addProvides(); nil != err {
			panic(err)
		}
		if err := application.setupStartup(); nil != err {
			panic(err)
		}
		// 增加内置的命令及参数
		if err := application.addInternalFlags(); nil != err {
			panic(err)
		}
		if err := application.addInternalCommands(); nil != err {
			panic(err)
		}
	})

	for _, opt := range opts {
		opt.apply(application.options)
	}

	return application
}

// Adds 添加各种组件到系统中
func (a *Application) Adds(components ...interface{}) (err error) {
	for _, component := range components {
		switch component.(type) {
		case app.Executor:
			err = a.AddExecutor(component.(app.Executor))
		case app.Serve:
			err = a.AddServes(component.(app.Serve))
		case app.Command:
			err = a.AddCommands(component.(app.Command))
		case app.Arg:
			err = a.AddArgs(component.(app.Arg))
		default:
			err = errors.New("不支持的类型")
		}

		if nil != err {
			break
		}
	}

	return
}

// AddServes 添加一个服务器到应用程序中
func (a *Application) AddServes(serves ...app.Serve) error {
	return a.Invoke(func(cmd *command.Serve) {
		for _, serve := range serves {
			cmd.AddServes(serve)
		}
	})
}

// AddCommands 添加一个可以被执行的命令到应用程序中
func (a *Application) AddCommands(commands ...app.Command) error {
	return a.Invoke(func(startup *cli.App) {
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
	return a.Invoke(func(startup *cli.App) {
		for _, argument := range args {
			parameter := argument
			startup.Flags = append(startup.Flags, parameter.ParseFlag())
		}
	})
}

// AddExecutor 添加执行器
func (a *Application) AddExecutor(executors ...app.Executor) (err error) {
	for _, executor := range executors {
		switch executor.Type() {
		case app.ExecutorTypeBeforeAll:
			a.beforeExecutors = append(a.beforeExecutors, executor)
		case app.ExecutorTypeAfterAll:
			a.afterExecutors = append(a.afterExecutors, executor)
		case app.ExecutorTypeBeforeServe:
			fallthrough
		case app.ExecutorTypeAfterServe:
			err = a.Invoke(func(serve *command.Serve) {
				serve.AddExecutors(executor)
			})
		}
	}

	return
}

// Provide 提供依赖关系
func (a *Application) Provide(constructor interface{}, opts ...provideOption) (err error) {
	options := defaultProvideOptions()
	for _, opt := range opts {
		opt.applyProvide(options)
	}

	return a.container.Provide(constructor)
}

// Provides 提供依赖关系
func (a *Application) Provides(constructors ...interface{}) (err error) {
	for _, constructor := range constructors {
		if err = a.container.Provide(constructor); nil != err {
			return
		}
	}

	return
}

// Invoke 获得依赖对象
func (a *Application) Invoke(function interface{}, opts ...invokeOption) error {
	options := defaultInvokeOptions()
	for _, opt := range opts {
		opt.applyInvoke(options)
	}

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

	// 添加启动器到依赖关系中
	if err = a.Provide(bootstrap); nil != err {
		return
	}
	// 加载用户启动器并做好配置
	if err = a.Invoke(func(bootstrap Bootstrap) error {
		return bootstrap.Setup()
	}); nil != err {
		return
	}

	// 执行生命周期方法
	if 0 != len(a.beforeExecutors) {
		if err = app.RunExecutors(a.beforeExecutors...); nil != err {
			return
		}
	}

	// 启动应用
	if err = a.Invoke(func(startup *cli.App) error {
		return startup.Run(os.Args)
	}); nil != err {
		return
	}

	// 执行生命周期方法
	if 0 != len(a.afterExecutors) {
		if err = app.RunExecutors(a.afterExecutors...); nil != err {
			return
		}
	}

	// 退出程序，解决最外层panic报错的问题
	// 原理：如果到这个地方还没有发生错误，程序正常退出，外层panic得不到执行
	// 如果发生错误，则所有代码都会返回error直到panic检测到，然后程序整体panic
	os.Exit(0)

	return
}

// LoadConfig 取得解析后的配置
func (a *Application) LoadConfig(config interface{}, opts ...option) (err error) {
	return a.config.Load(config, opts...)
}

func (a *Application) setupStartup() error {
	cli.AppHelpTemplate = a.options.helpAppTemplate
	cli.CommandHelpTemplate = a.options.helpCommandTemplate
	cli.SubcommandHelpTemplate = a.options.helpSubcommandTemplate
	cli.VersionPrinter = func(ctx *cli.Context) {
		_ = a.Invoke(func(version *command.Version) error {
			return version.Run(app.NewContext(ctx))
		})
	}

	return a.Invoke(func(startup *cli.App) {
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

		Serve   *command.Serve
		Version *command.Version
	}

	return a.Invoke(func(in commandsIn) error {
		return a.AddCommands(in.Serve, in.Version)
	})
}

func (a *Application) addInternalFlags() error {
	return a.Invoke(func(startup *cli.App) {
		startup.Flags = append(startup.Flags, &cli.StringFlag{
			Name:        "conf",
			Aliases:     []string{"c"},
			Usage:       "指定配置文件路径",
			Value:       "./conf/application.yaml",
			DefaultText: "./conf/application.yaml",
		})
	})
}

func (a *Application) addProvides() (err error) {
	if err = a.Provides(gox.NewSnowflake); nil != err {
		return
	}
	if err = a.Provides(command.NewServe, command.NewVersion); nil != err {
		return
	}
	if err = a.Provides(appName, appVersion, buildVersion, buildTime, scmRevision, scmBranch, goVersion); nil != err {
		return
	}
	if err = a.Provides(newApp, app.NewDefaultService); nil != err {
		return
	}

	// 注入配置
	if err = a.Provide(func() *Config {
		return a.config
	}); nil != err {
		return
	}
	// 注入日志
	if err = a.Provide(func() app.Logger {
		return a.options.logger
	}); nil != err {
		return
	}
	// 注入自身
	if err = a.Provide(func() *Application {
		return a
	}); nil != err {
		return
	}

	return
}
