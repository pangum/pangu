package core

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"

	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/harluo/boot/internal"
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/boot/internal/core/internal/command"
	"github.com/harluo/boot/internal/core/internal/core"
	"github.com/harluo/boot/internal/core/internal/get"
	"github.com/harluo/boot/internal/core/internal/message"
	"github.com/harluo/boot/internal/internal/config"
	"github.com/harluo/boot/internal/internal/constant"
	"github.com/harluo/boot/internal/internal/kernel"
	"github.com/harluo/boot/internal/runtime"
	"github.com/harluo/di"
	"github.com/urfave/cli/v2"
)

var shadow *Application

type Application struct {
	params    *config.Application
	container *di.Container

	stoppers []application.Stopper
	befores  []application.Before
	afters   []application.After

	logger log.Logger
}

func New(param *config.Application) (application *Application) {
	once.Do(create(param))
	application = shadow

	return
}

func create(params *config.Application) func() {
	return func() {
		shadow = new(Application)
		shadow.params = params
		shadow.container = di.New().Validate().Get()

		shadow.befores = make([]application.Before, 0)
		shadow.afters = make([]application.After, 0)
		shadow.stoppers = make([]application.Stopper, 0)

		// ! 这个操作必须在创建的时候就执行，因为在后续插件启动时会寻找下面的依赖，而在这个时候启动方法还没有执行
		shadow.container.Dependency().Puts(
			shadow.putSelf, // 注入自身
		).Invalidate().Build().Apply()
	}
}

// Run 启动应用程序
func (a *Application) Run(constructor runtime.Constructor) {
	var err error
	defer a.finally(&err)

	dependency := a.container.Dependency()
	if cae := a.createShell(); nil != cae { // 创建运行壳
		err = cae
	} else if ape := a.addDependency(constructor); nil != ape { // 添加应用程序本身包含依赖，供后续应用启动执行
		err = ape
	} else {
		err = dependency.Get(a.boot).Build().Build().Inject()
	}
}

// Add 添加各种组件到系统中
func (a *Application) Add(required application.Command, optionals ...application.Command) (err error) {
	for _, _command := range append([]application.Command{required}, optionals...) {
		err = a.addCommand(_command)
		if nil != err {
			break
		}
	}

	return
}

func (a *Application) finally(err *error) {
	if finally := recover(); nil != finally {
		fmt.Println(gox.Stack(constant.ApplicationStacktrace, constant.ApplicationSkip))
		os.Exit(a.params.Code.Panic)
	}

	if nil == *err {
		os.Exit(a.params.Code.Success)
	} else {
		fmt.Println((*err).Error())
		os.Exit(a.params.Code.Failed)
	}
}

func (a *Application) addCommand(command application.Command) error {
	return a.container.Dependency().Get(func(shell *core.Shell) {
		shell.Commands = append(shell.Commands, a.convertCommand(command))
	}).Build().Build().Inject()
}

func (a *Application) convertCommand(command application.Command) (appended *cli.Command) {
	appended = new(cli.Command)
	appended.Name = command.Name()
	appended.Action = a.action(command)

	// 通过反射设置非必须参数，减少编码成本
	typer := core.NewTyper(command)
	appended.Aliases = typer.Aliases()
	appended.Usage = typer.Usage()
	appended.Description = typer.Description()
	appended.Category = typer.Category()
	appended.Hidden = typer.Hidden()
	appended.Action = a.action(command)
	if converted, ok := command.(kernel.Subcommands); ok {
		for _, subcommand := range converted.Subcommands() {
			appended.Subcommands = append(appended.Subcommands, a.convertCommand(subcommand))
		}
	}
	if converted, ok := command.(kernel.Arguments); ok {
		for _, argument := range converted.Arguments() {
			appended.Flags = append(appended.Flags, core.NewArgument(argument).Flag())
		}
	}

	// 生命周期方法
	if converted, ok := command.(application.Before); ok {
		a.befores = append(a.befores, converted)
	}
	if converted, ok := command.(application.After); ok {
		a.afters = append(a.afters, converted)
	}
	if converted, ok := command.(application.Stopper); ok {
		a.stoppers = append(a.stoppers, converted)
	}

	return
}

func (a *Application) addArguments(get get.Arguments) error {
	return a.container.Dependency().Get(func(shell *core.Shell) {
		for _, argument := range get.Arguments {
			shell.Flags = append(shell.Flags, core.NewArgument(argument).Flag())
		}
	}).Build().Build().Inject()
}

func (a *Application) boot() (err error) {
	// 优雅退出
	go a.graceful(&err)

	canceled, cancel := context.WithTimeout(context.Background(), a.params.Timeout.Startup)
	defer cancel()

	dependency := a.container.Dependency()
	if aae := dependency.Get(a.addArguments).Build().Build().Inject(); nil != aae { // 增加参数
		err = aae
	} else if ace := dependency.Get(a.addCommands).Build().Build().Inject(); nil != ace { // 增加命令
		err = ace
	} else if re := dependency.Get(a.run).Build().Build().Inject(); nil != re { // 执行整个程序
		err = re
	} else if bpe := a.params.Banner.Print(); nil != bpe { // 打印标志信息
		err = bpe
	} else if ie := dependency.Get(a.initialize(canceled)).Build().Build().Inject(); nil != ie { // 执行初始化方法
		err = ie
	} else if sle := dependency.Get(a.setLogger).Build().Build().Inject(); nil != sle { // 为执行壳设置日志器
		err = sle
	} else if ple := a.putLogger(); nil != ple {
		err = ple
	} else if se := dependency.Get(a.startup(canceled)).Build().Build().Inject(); nil != se { // 执行初始化方法
		err = se
	}

	return
}

func (a *Application) initialize(ctx context.Context) func(get.Initializers) error {
	return func(get get.Initializers) (err error) {
		for _, initializer := range get.Initializers {
			err = initializer.Initialize(ctx)
			if nil != err {
				break
			}
		}

		return
	}
}

func (a *Application) startup(ctx context.Context) func(Starter) error {
	return func(starter Starter) (err error) {
		if bse := a.beforeStater(ctx, starter); nil != bse { // 生命周期方法
			err = bse
		} else if se := starter.Startup(a); nil != se { // 加载用户启动器并做好配置
			err = se
		} else if cbe := a.before(ctx); nil != cbe { // 执行命令生命周期方法
			err = cbe
		} else if ple := a.putLogger(); nil != ple { // 日志
			err = ple
		} else if cae := a.after(ctx); nil != cae { // 执行命令生命周期方法
			err = cae
		} else if ase := a.afterStater(ctx, starter); nil != ase { // 执行命令生命周期方法
			err = ase
		}

		return
	}
}

func (a *Application) run(shell *core.Shell) error {
	return shell.Run(os.Args)
}

func (a *Application) createShell() (err error) {
	cli.AppHelpTemplate = a.params.Help.App
	cli.CommandHelpTemplate = a.params.Help.Command
	cli.SubcommandHelpTemplate = a.params.Help.Subcommand
	cli.VersionPrinter = a.versionPrinter

	// 定制版本标志
	version := new(cli.BoolFlag)
	version.Name = "version"
	version.Aliases = []string{
		"v",
		"ver",
	}
	version.Usage = "显示应用程序版本信息"
	cli.VersionFlag = version

	// 定制帮助信息
	help := new(cli.BoolFlag)
	help.Name = "help"
	help.Aliases = []string{
		"h",
	}
	help.Usage = "显示所有命令或者帮助信息"
	cli.HelpFlag = help

	// 配置应用
	err = a.container.Dependency().Get(a.setupShell).Build().Build().Inject()

	return
}

func (a *Application) addCommands(get get.Commands) (err error) {
	for _, _command := range get.Commands {
		err = a.addCommand(_command)
		if nil != err {
			break
		}
	}

	return
}

func (a *Application) addDependency(constructor runtime.Constructor) (err error) {
	dependency := a.container.Dependency().Invalidate()
	if ve := a.verify(constructor); nil != ve {
		err = ve
	} else if pce := dependency.Put(constructor).Build().Build().Inject(); nil != pce {
		err = pce
	}

	return
}

func (a *Application) setupShell(shell *core.Shell) {
	shell.Name = internal.Name
	shell.Description = a.params.Description
	shell.Usage = a.params.Usage
	shell.Copyright = a.params.Copyright
	shell.Metadata = a.params.Metadata
	shell.Authors = a.params.Authors.Cli()
	shell.AllowExtFlags = true
}

func (a *Application) versionPrinter(ctx *cli.Context) {
	_ = a.container.Dependency().Get(func(info *command.Info) error {
		return info.Run(runtime.NewContext(ctx))
	}).Build().Build().Inject()
}

func (a *Application) putSelf() *Application {
	return a
}

func (a *Application) putLogger() (err error) {
	dependency := a.container.Dependency()
	if ge := dependency.Get(a.getLogger).Build().Build().Inject(); nil != ge {
		err = dependency.Put(a.supplyLogger).Build().Build().Inject() // !当出错或未成功设置时，重置日志器
	}

	return
}

func (a *Application) getLogger(logger get.Logger) {
	if nil != logger.Optional { // !只有在确实有外部日志器的情况下才允许覆盖
		a.logger = logger.Optional
	}
}

func (a *Application) getDefaultLogger() (logger log.Logger) {
	logger = log.New().Apply()
	if level, ok := os.LookupEnv(constant.EnvironmentLoggingLevel); ok {
		logger.Enable(log.ParseLevel(level))
	}

	return
}

func (a *Application) supplyLogger() log.Logger {
	return a.logger
}

func (a *Application) verify(bootstrap runtime.Constructor) (err error) {
	if a.params.Validate {
		return
	}

	typ := reflect.TypeOf(bootstrap)
	if reflect.Func != typ.Kind() { // 构造方法必须是方法不能是其它类型
		err = exception.New().Message(message.ConstructorMustFunc).Field(field.New("bootstrap", typ.String())).Build()
	} else if 0 == typ.NumIn() { // 构造方法必须有依赖项
		name := runtime.FuncForPC(reflect.ValueOf(bootstrap).Pointer()).Name()
		err = exception.New().Message(message.BootstrapMustHasDependencies).Field(field.New("bootstrap", name)).Build()
	} else if 1 != typ.NumOut() || reflect.TypeOf((*Starter)(nil)).Elem() != typ.Out(constant.IndexFirst) {
		// 只能返回一个类型为Bootstrap返回值
		err = exception.New().Message(message.BootstrapMustReturnBootstrap).Build()
	}

	return
}

func (a *Application) action(command application.Command) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return command.Run(runtime.NewContext(ctx))
	}
}

func (a *Application) graceful(err *error) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	current := <-signals
	a.logger.Info("收到系统信号", field.New("signal", current))

	canceled, cancel := context.WithTimeout(context.Background(), a.params.Timeout.Quit)
	defer cancel()
	for _, stopper := range a.stoppers {
		*err = stopper.Stop(canceled)
		if nil != *err {
			break
		}
	}
}

func (a *Application) setLogger(shell *core.Shell) {
	shell.Logger(a.logger)
}

func (a *Application) beforeStater(ctx context.Context, stater Starter) (err error) {
	if converted, ok := stater.(application.Before); ok {
		err = converted.Before(ctx)
	}

	return
}

func (a *Application) before(ctx context.Context) (err error) {
	a.logger.Info(
		"准备启动应用",
		field.New(strings.ToLower(constant.ColumnName), internal.GetName()),
		field.New(strings.ToLower(constant.ColumnVersion), internal.GetVersion()),
		field.New(strings.ToLower(constant.ColumnBuild), internal.GetBuild()),
		field.New(strings.ToLower(constant.ColumnComplied), internal.GetCompiled()),
		field.New(strings.ToLower(constant.ColumnRevision), internal.GetRevision()),
		field.New(strings.ToLower(constant.ColumnBranch), internal.GetBranch()),
		field.New(strings.ToLower(constant.ColumnRuntime), internal.Runtime),
	)
	for _, before := range a.befores {
		err = before.Before(ctx)
		if nil != err {
			break
		}
	}

	return
}

func (a *Application) after(ctx context.Context) (err error) {
	for _, after := range a.afters {
		err = after.After(ctx)
		if nil != err {
			break
		}
	}

	return
}

func (a *Application) afterStater(ctx context.Context, stater Starter) (err error) {
	if converted, ok := stater.(application.After); ok {
		err = converted.After(ctx)
	}

	return
}
