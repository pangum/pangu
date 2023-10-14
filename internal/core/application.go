package core

import (
	"os"
	"sync"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/command"
	"github.com/pangum/pangu/internal/get"
	"github.com/pangum/pangu/internal/internal/builder"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/storezhang/dig"
	"github.com/urfave/cli/v2"
)

var (
	application *Application
	once        sync.Once
)

type Application struct {
	config    *Config
	param     *param.Application
	container *dig.Container
	// 影子启动器，用来处理额外的命令或者参数，因为正常的启动器无法完成此操作，原因是
	// 正常的启动器只提供一个Run方法来处理传数的参数，而此方法一旦执行，就意味着内部的命令也开始执行，而此时依赖关系还没有准备好
	shadow *runtime.Shadow

	verify *Verify
}

func New(param *param.Application) *Application {
	once.Do(creator(param))
	// 可以在每次调用时，修改部分配置
	application.param.Override(param)

	return application
}

func creator(param *param.Application) func() {
	return func() {
		create(param)
	}
}

func create(param *param.Application) {
	application = new(Application)
	application.param = param
	application.container = dig.New()
	application.shadow = runtime.NewShadow()
	// 注入配置对象，后续使用
	// TODO app.config = pangu.newConfig(_options.configOptions)
}

func (a *Application) Dependency() *builder.Dependency {
	return builder.NewDependency(a.container)
}

func (a *Application) Dependencies() *builder.Dependencies {
	return builder.NewDependencies(a.container)
}

// Run 启动应用程序
func (a *Application) Run(bootstrap Bootstrap) {
	var err error
	defer a.finally(&err)

	dependency := a.Dependency().Build()
	if ese := a.param.Environments.Set(); nil != ese { // 添加环境变量
		err = ese
	} else if ape := a.addDependency(); nil != ape { // 添加内置的依赖
		err = ape
	} else if cae := a.createApp(); nil != cae { // 创建应用
		err = cae
	} else if ace := dependency.Get(a.addInternalCommands); nil != ace { // 增加内置的命令及参数
		err = ace
	} else if bpe := a.param.Banner.Print(); nil != bpe { // 输出标志信息
		err = bpe
	} else if bde := dependency.Get(a.bind); nil != bde { // 绑定参数和命令到内部变量或者命令上
		err = bde
	} else if bse := bootstrap.Startup(); nil != bse { // 加载用户启动器并做好配置
		err = bse
	} else if rbe := bootstrap.Before(); nil != rbe { // 执行生命周期方法
		err = rbe
	} else if re := dependency.Get(a.run); nil != re { // 执行整个程序
		err = re
	} else if rae := bootstrap.After(); nil != rae { // 执行生命周期方法
		err = rae
	}

	return
}

/* TODO // Load 取得解析后的配置
func (a *Application) Load(config any, opts ...pangu.configOption) (err error) {
	return a.config.Load(config, opts...)
}

// Watch 监控配置变化
func (a *Application) Watch(config any, watcher config.configWatcher) (err error) {
	return gfx.Watch(a.config.path, pangu.newConfigFileWatcher(config, a.config.path, watcher, a.config.param))
}*/

// Add 添加各种组件到系统中
func (a *Application) Add(components ...any) (err error) {
	for _, component := range components {
		switch typ := component.(type) {
		case app.Serve:
			err = a.addServe(typ)
		case app.Command:
			err = a.addCommands(typ)
		case app.Argument:
			err = a.addArgs(typ)
		default:
			err = exc.NewField("不支持的类型", field.New("type", typ))
		}

		if nil != err {
			break
		}
	}

	return
}

func (a *Application) finally(err *error) {
	if finally := recover(); nil != finally {
		panic(finally)
		// TODO os.Exit(internal.ExitCodeFailed)
	}

	if nil == *err {
		// TODO os.Exit(internal.ExitCodeOk)
	} else {
		// TODO os.Exit(internal.ExitCodeFailed)
	}
}

func (a *Application) addServe(serves ...app.Serve) error {
	return a.Dependency().Build().Get(func(cmd *command.Serve) {
		for _, serve := range serves {
			cmd.Add(serve)
		}
	})
}

func (a *Application) addCommands(commands ...app.Command) error {
	return a.Dependency().Build().Get(func(shell *runtime.Shell) {
		shell.Commands = append(shell.Commands, app.Commands(commands).Cli()...)
	})
}

func (a *Application) addArgs(args ...app.Argument) error {
	return a.Dependency().Build().Get(func(shell *runtime.Shell) {
		for _, arg := range args {
			_arg := arg
			shell.Flags = append(shell.Flags, _arg.Flag())
		}
	})
}

func (a *Application) bind(shell *runtime.Shell) error {
	a.config.bind(shell, a.shadow)

	// 影子执行，将所有参数都绑定到具体的变量上或者命令上
	return a.shadow.Run(a.args())
}

func (a *Application) run(shell *cli.App) error {
	return shell.Run(a.args())
}

func (a *Application) args() []string {
	return os.Args
}

func (a *Application) createApp() (err error) {
	cli.AppHelpTemplate = a.param.App
	cli.CommandHelpTemplate = a.param.Command
	cli.SubcommandHelpTemplate = a.param.Subcommand
	cli.VersionPrinter = a.versionPrinter
	err = a.Dependency().Build().Get(a.setupApp)

	return
}

func (a *Application) addInternalCommands(get get.Command) error {
	return a.addCommands(get.Serve, get.Info, get.Version)
}

func (a *Application) addDependency() (err error) {
	dependency := a.Dependency().Build()
	dependencies := a.Dependencies().Build()
	if cde := dependencies.Put(command.All...); nil != cde { // 注入所有内置命令
		err = cde
	} else if nse := dependency.Put(runtime.NewShell); nil != nse { // 注入运行壳
		err = nse
	} else if ce := dependency.Put(a.putConfig); nil != ce { // 注入配置
		err = ce
	} else if se := dependency.Put(a.putSelf); nil != se { // 注入自身
		err = se
	}

	return
}

func (a *Application) setupApp(app *cli.App) {
	app.Name = runtime.Name
	app.Description = a.param.Description
	app.Usage = a.param.Usage
	app.Copyright = a.param.Copyright
	app.Metadata = a.param.Metadata
	app.Authors = a.param.Authors.Cli()
}

func (a *Application) versionPrinter(ctx *cli.Context) {
	_ = a.Dependency().Build().Get(func(info *command.Info) error {
		return info.Run(app.NewContext(ctx))
	})
}

func (a *Application) putConfig() *Config {
	return a.config
}

func (a *Application) putSelf() *Application {
	return a
}
