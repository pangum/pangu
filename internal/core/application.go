package core

import (
	"fmt"
	"os"
	"sync"

	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/command"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/get"
	"github.com/pangum/pangu/internal/internal/builder"
	"github.com/pangum/pangu/internal/internal/verifier"
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
	params    *param.Application
	container *dig.Container
	// 影子启动器，用来处理额外的命令或者参数，因为正常的启动器无法完成此操作，原因是
	// 正常的启动器只提供一个Run方法来处理传数的参数，而此方法一旦执行，就意味着内部的命令也开始执行，而此时依赖关系还没有准备好
	shadow *runtime.Shadow
}

func New(param *param.Application) *Application {
	once.Do(creator(param))
	// 可以在每次调用时，修改部分配置
	application.params.Override(param)

	return application
}

func creator(param *param.Application) func() {
	return func() {
		create(param)
	}
}

func create(params *param.Application) {
	application = new(Application)
	application.params = params
	application.container = dig.New()
	application.shadow = runtime.NewShadow()
	application.config = NewConfig(params.Config)
}

func (a *Application) Dependency() *builder.Dependency {
	return builder.NewDependency(a.container, a.params)
}

func (a *Application) Dependencies() *builder.Dependencies {
	return builder.NewDependencies(a.container, a.params)
}

// Run 启动应用程序
func (a *Application) Run(constructor runtime.Constructor) {
	var err error
	defer a.finally(&err)

	dependency := a.Dependency().Build()
	if ese := a.params.Environments.Set(); nil != ese { // 添加环境变量
		err = ese
	} else if ape := a.addDependency(constructor); nil != ape { // 添加内置的依赖
		err = ape
	} else if cae := a.createApp(); nil != cae { // 创建应用
		err = cae
	} else if ace := dependency.Get(a.addCommands); nil != ace { // 增加内置的命令及参数
		err = ace
	} else if bpe := a.params.Banner.Print(); nil != bpe { // 输出标志信息
		err = bpe
	} else if bde := dependency.Get(a.bind); nil != bde { // 绑定参数和命令到内部变量或者命令上
		err = bde
	} else {
		err = dependency.Get(a.boot)
	}

	return
}

// Add 添加各种组件到系统中
func (a *Application) Add(components ...any) (err error) {
	for _, component := range components {
		switch converted := component.(type) {
		case app.Serve:
			err = a.addServe(converted)
		case app.Command:
			err = a.addCommand(converted)
		case app.Argument:
			err = a.addArg(converted)
		default:
			err = exc.NewField("不支持的类型", field.New("type", converted))
		}

		if nil != err {
			break
		}
	}

	return
}

func (a *Application) finally(err *error) {
	if finally := recover(); nil != finally {
		fmt.Println(gox.Stack(constant.ApplicationStacktrace, constant.ApplicationSkip))
		os.Exit(a.params.Panic)
	}

	if nil == *err {
		os.Exit(a.params.Success)
	} else {
		fmt.Println((*err).Error())
		os.Exit(a.params.Failed)
	}
}

func (a *Application) addServe(serves ...app.Serve) error {
	return a.Dependency().Build().Get(func(cmd *command.Serve) {
		for _, serve := range serves {
			cmd.Add(serve)
		}
	})
}

func (a *Application) addCommand(commands ...app.Command) error {
	return a.Dependency().Build().Get(func(shell *runtime.Shell) {
		shell.Commands = append(shell.Commands, app.Commands(commands).Cli()...)
	})
}

func (a *Application) addArg(args ...app.Argument) error {
	return a.Dependency().Build().Get(func(shell *runtime.Shell) {
		for _, arg := range args {
			_arg := arg
			shell.Flags = append(shell.Flags, _arg.Flag())
		}
	})
}

func (a *Application) bind(shell *runtime.Shell) (err error) {
	// 接收配置文件路径参数
	a.config.bind(shell, a.shadow)
	// 影子执行，只有这样才能正确的使用配置文件的路径参数
	err = a.shadow.Run(a.args())

	return
}

func (a *Application) boot(bootstrap runtime.Bootstrap) (err error) {
	dependency := a.Dependency().Build()
	if bse := bootstrap.Startup(); nil != bse { // 加载用户启动器并做好配置
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

func (a *Application) run(shell *runtime.Shell) error {
	return shell.Run(a.args())
}

func (a *Application) args() []string {
	return os.Args
}

func (a *Application) createApp() (err error) {
	cli.AppHelpTemplate = a.params.App
	cli.CommandHelpTemplate = a.params.Command
	cli.SubcommandHelpTemplate = a.params.Subcommand
	cli.VersionPrinter = a.versionPrinter
	err = a.Dependency().Build().Get(a.setupApp)

	return
}

func (a *Application) addCommands(get get.Command) error {
	return a.addCommand(get.Serve, get.Info, get.Version)
}

func (a *Application) addDependency(constructor runtime.Constructor) (err error) {
	dependency := a.Dependency().Build()
	dependencies := a.Dependencies().Build()
	boostrap := verifier.NewBoostrap(a.params)
	if ve := boostrap.Verify(constructor); nil != ve {
		err = ve
	} else if pce := dependency.Put(constructor); nil != pce {
		err = pce
	} else if ple := dependency.Put(a.putLogger); nil != ple {
		err = ple
	} else if pse := dependencies.Put(command.NewServe); nil != pse { // 注入服务命令
		err = pse
	} else if pie := dependencies.Put(command.NewInfo); nil != pie { // 注入信息命令
		err = pie
	} else if pve := dependencies.Put(command.NewVersion); nil != pve { // 注入版本命令
		err = pve
	} else if nse := dependency.Put(runtime.NewShell); nil != nse { // 注入运行壳
		err = nse
	} else if ce := dependency.Put(a.putConfig); nil != ce { // 注入配置
		err = ce
	} else if se := dependency.Put(a.putSelf); nil != se { // 注入自身
		err = se
	}

	return
}

func (a *Application) setupApp(shell *runtime.Shell) {
	shell.Name = runtime.Name
	shell.Description = a.params.Description
	shell.Usage = a.params.Usage
	shell.Copyright = a.params.Copyright
	shell.Metadata = a.params.Metadata
	shell.Authors = a.params.Authors.Cli()
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

func (a *Application) putLogger() (logger app.Logger) {
	err := a.Dependency().Build().Get(func(dependency simaqian.Logger) {
		logger = dependency
	})
	if nil != err {
		logger = simaqian.Default()
	}

	return
}
