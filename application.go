package pangu

import (
	"os"
	"reflect"
	"runtime"
	"sync"

	"github.com/goexl/exc"
	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/dig"
	"github.com/urfave/cli/v2"
)

// Application 应用程序，可以加入两种种类型的程序
// Serve 用于描述应用程序内的服务
// Command 用于描述应用程序内可以被执行的命令
type Application struct {
	config    *Config
	options   *options
	container *dig.Container
	// 影子启动器，用来处理额外的命令或者参数，因为正常的启动器无法完成此操作，原因是
	// 正常的启动器只提供一个Run方法来处理传数的参数，而此方法一旦执行，就意味着内部的命令也开始执行，而此时依赖关系还没有准备好
	shadow *cli.App

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
		_options := defaultOptions()
		application = &Application{
			options:   _options,
			container: dig.New(),
			shadow:    newShadow(),

			beforeExecutors: make([]app.Executor, 0),
			afterExecutors:  make([]app.Executor, 0),
		}
		// 注入配置对象，后续使用
		application.config = newConfig(_options.configOptions)

		// 初始化内置变量及内置命令
		// 之所以在这儿完全初始化，是因为可能会在任意地方调用方法，而在调用这些方法时要求必须完成初始化
		// 从另外一个角度上来说，如果初始化失败，panic是被允许的（盘古是应用程序的基础，基础初始化有错误，不panic更待何时）
		// Adds
		// AddCommands
		// AddExecutors
		// AddArgs
		// AddFlags
		if err := application.addEnvs(); nil != err {
			panic(err)
		}
		if err := application.addProvides(); nil != err {
			panic(err)
		}
		if err := application.setupStartup(); nil != err {
			panic(err)
		}
		// 增加内置的命令及参数
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
func (a *Application) Adds(components ...any) (err error) {
	for _, component := range components {
		switch typ := component.(type) {
		case app.Executor:
			err = a.AddExecutor(typ)
		case app.Serve:
			err = a.AddServes(typ)
		case app.Command:
			err = a.AddCommands(typ)
		case app.Argument:
			err = a.AddArgs(typ)
		default:
			err = exc.NewField(`不支持的类型`, field.New("type", typ))
		}

		if nil != err {
			break
		}
	}

	return
}

// AddServes 添加一个服务器到应用程序中
func (a *Application) AddServes(serves ...app.Serve) error {
	return a.Invoke(func(cmd *cmd.Serve) {
		for _, serve := range serves {
			cmd.AddServes(serve)
		}
	})
}

// AddCommands 添加一个可以被执行的命令到应用程序中
func (a *Application) AddCommands(commands ...app.Command) error {
	return a.Invoke(func(shell *cli.App) {
		shell.Commands = append(shell.Commands, a.commands(commands...)...)
	})
}

// AddArgs 添加参数
func (a *Application) AddArgs(args ...app.Argument) error {
	return a.Invoke(func(shell *cli.App) {
		for _, arg := range args {
			_arg := arg
			shell.Flags = append(shell.Flags, _arg.Flag())
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
			err = a.Invoke(func(serve *cmd.Serve) {
				serve.AddExecutors(executor)
			})
		}
	}

	return
}

// Provide 提供依赖关系
func (a *Application) Provide(constructor any, opts ...provideOption) (err error) {
	_options := defaultProvideOptions()
	for _, opt := range opts {
		opt.applyProvide(_options)
	}

	// 验证构造方法是否合法
	if err = a.validateConstructor(constructor); nil != err {
		return
	}
	err = a.container.Provide(constructor)

	return
}

// Dependency 提供依赖关系，如果依赖关系有错，退出
// Provide 方法的特殊封装，如果有错误直接退出
func (a *Application) Dependency(constructor any, opts ...provideOption) {
	if err := a.Provide(constructor, opts...); nil != err {
		panic(err)
	}
}

// Provides 提供依赖关系
func (a *Application) Provides(constructors ...any) (err error) {
	for _, constructor := range constructors {
		if err = a.Provide(constructor); nil != err {
			return
		}
	}

	return
}

// Dependencies 提供依赖关系，如果依赖关系有错，退出
// Provides 方法的特殊封装，如果有错误直接退出
func (a *Application) Dependencies(constructors ...any) {
	for _, constructor := range constructors {
		if err := a.Provide(constructor); nil != err {
			panic(err)
		}
	}
}

// Invoke 获得依赖对象
func (a *Application) Invoke(function any, opts ...invokeOption) error {
	_options := defaultInvokeOptions()
	for _, opt := range opts {
		opt.applyInvoke(_options)
	}

	return a.container.Invoke(function)
}

// Run 启动应用程序
func (a *Application) Run(constructor any) (err error) {
	// 验证启动器构造方法是否合法
	if err = a.validateBootstrap(constructor); nil != err {
		return
	}

	// 输出标志信息
	if `` != a.options.banner.data {
		if err = a.options.banner.print(); nil != err {
			return
		}
	}

	// 添加启动器到依赖关系中
	if err = a.Provide(constructor); nil != err {
		return
	}
	// 绑定参数和命令到内部变量或者命令上
	if err = a.Invoke(a.bind); nil != err {
		return
	}
	// 加载用户启动器并做好配置
	if err = a.Invoke(a.bootstrap); nil != err {
		return
	}

	// 执行生命周期方法
	if 0 != len(a.beforeExecutors) {
		if err = app.RunExecutors(a.beforeExecutors...); nil != err {
			return
		}
	}

	// 执行整个程序
	if err = a.Invoke(a.run); nil != err {
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

// Load 取得解析后的配置
func (a *Application) Load(config any, opts ...configOption) (err error) {
	return a.config.Load(config, opts...)
}

// Watch 监控配置变化
func (a *Application) Watch(config any, watcher configWatcher) (err error) {
	return gfx.Watch(a.config.path, newConfigFileWatcher(config, a.config.path, watcher, a.config.options))
}

func (a *Application) bind(shell *cli.App) error {
	a.config.bind(shell, a.shadow)

	// 影子执行，将所有参数都绑定到具体的变量上或者命令上
	return a.shadow.Run(a.args())
}

func (a *Application) bootstrap(bootstrap Bootstrap) error {
	return bootstrap.Startup()
}

func (a *Application) run(shell *cli.App) error {
	return shell.Run(a.args())
}

func (a *Application) args() []string {
	return os.Args
}

func (a *Application) commands(acs ...app.Command) (commands []*cli.Command) {
	if 0 != len(acs) {
		commands = make([]*cli.Command, 0, len(acs))
	}

	for _, ac := range acs {
		command := ac
		commands = append(commands, &cli.Command{
			Name:        command.Name(),
			Aliases:     command.Aliases(),
			Usage:       command.Usage(),
			Description: command.Description(),
			Subcommands: a.commands(command.Subcommands()...),
			Category:    command.Category(),
			Flags:       a.flags(command.Arguments()...),
			Hidden:      command.Hidden(),
			Action: func(ctx *cli.Context) error {
				return command.Run(app.NewContext(ctx))
			},
		})
	}

	return
}

func (a *Application) flags(ins ...app.Argument) (flags []cli.Flag) {
	if 0 != len(ins) {
		flags = make([]cli.Flag, 0, len(ins))
	}

	for _, in := range ins {
		flags = append(flags, in.Flag())
	}

	return
}

func (a *Application) validateBootstrap(constructor any) (err error) {
	if a.options.verify {
		return
	}

	constructorType := reflect.TypeOf(constructor)
	// 构造方法必须是方法不能是其它类型
	if reflect.Func != constructorType.Kind() {
		err = exc.NewField(exceptionConstructorMustFunc, field.New("constructor", constructorType.String()))
	}
	if nil != err {
		return
	}

	// 构造方法必须有依赖项
	if 0 == constructorType.NumIn() {
		constructorName := runtime.FuncForPC(reflect.ValueOf(constructor).Pointer()).Name()
		err = exc.NewField(exceptionBootstrapMustHasDependencies, field.New("constructor", constructorName))
	}
	if nil != err {
		return
	}

	// 只能返回一个类型为Bootstrap返回值
	returnsCount := constructorType.NumOut()
	if 1 != returnsCount || reflect.TypeOf((*Bootstrap)(nil)).Elem() != constructorType.Out(firstIndex) {
		err = exc.NewMessage(exceptionBootstrapMustReturnBootstrap)
	}

	return
}

func (a *Application) validateConstructor(constructor any) (err error) {
	if a.options.verify {
		return
	}

	constructorType := reflect.TypeOf(constructor)
	// 构造方法必须是方法不能是其它类型
	if reflect.Func != constructorType.Kind() {
		err = exc.NewField(exceptionConstructorMustFunc, field.New("constructor", constructorType.String()))
	}
	if nil != err {
		return
	}

	// 构造方法必须有返回值
	if 0 == constructorType.NumOut() {
		constructorName := runtime.FuncForPC(reflect.ValueOf(constructor).Pointer()).Name()
		err = exc.NewField(exceptionConstructorMustReturn, field.New("constructor", constructorName))
	}

	return
}

func (a *Application) setupStartup() error {
	cli.AppHelpTemplate = a.options.helpAppTemplate
	cli.CommandHelpTemplate = a.options.helpCommandTemplate
	cli.SubcommandHelpTemplate = a.options.helpSubcommandTemplate
	cli.VersionPrinter = func(ctx *cli.Context) {
		_ = a.Invoke(func(info *cmd.Info) error {
			return info.Run(app.NewContext(ctx))
		})
	}

	return a.Invoke(func(shell *cli.App) {
		shell.Name = Name
		shell.Description = a.options.description
		shell.Usage = a.options.usage
		shell.Copyright = a.options.copyright
		shell.Metadata = a.options.metadata
		if 0 != len(a.options.authors) {
			authors := make([]*cli.Author, 0, len(a.options.authors))
			for _, _author := range a.options.authors {
				authors = append(authors, &cli.Author{
					Name:  _author.name,
					Email: _author.email,
				})
			}
			shell.Authors = authors
		}
	})
}

func (a *Application) addEnvs() (err error) {
	if 0 == len(a.options.environments) {
		return
	}

	for _, env := range a.options.environments {
		if err = os.Setenv(env.key, env.value); nil != err {
			return
		}
	}

	return
}

func (a *Application) addInternalCommands() error {
	type commandIn struct {
		In

		Serve   *cmd.Serve
		Info    *cmd.Info
		Version *cmd.Version
	}

	return a.Invoke(func(in commandIn) error {
		return a.AddCommands(in.Serve, in.Info, in.Version)
	})
}

func (a *Application) addProvides() (err error) {
	if err = a.Provides(cmd.NewServe, cmd.NewInfo, cmd.NewVersion); nil != err {
		return
	}
	if err = a.Provides(name, version, build, timestamp, revision, branch, golang); nil != err {
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
