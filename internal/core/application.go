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
	"github.com/harluo/boot/internal/internal/checker"
	"github.com/harluo/boot/internal/internal/config"
	"github.com/harluo/boot/internal/internal/constant"
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
		shadow.container = di.New().Validate().Instance()

		shadow.befores = make([]application.Before, 0)
		shadow.afters = make([]application.After, 0)
		shadow.stoppers = make([]application.Stopper, 0)

		// ! 这个操作必须在创建的时候就执行，因为在后续插件启动时会寻找下面的依赖，而在这个时候启动方法还没有执行
		shadow.container.Put(
			shadow.putSelf, // 注入自身
		).Invalidate().Build().Apply()
	}
}

// Run 启动应用程序
func (a *Application) Run(constructor runtime.Constructor) {
	var err error
	defer a.finally(&err)

	if cae := a.createShell(); nil != cae { // 创建运行壳
		err = cae
	} else if ape := a.addDependency(constructor); nil != ape { // 添加应用程序本身包含依赖，供后续应用启动执行
		err = ape
	} else {
		err = a.container.Get(a.boot).Build().Inject()
	}
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
	for _, subcommand := range typer.Subcommands() {
		appended.Subcommands = append(appended.Subcommands, a.convertCommand(subcommand))
	}
	for _, argument := range typer.Arguments() {
		appended.Flags = append(appended.Flags, core.NewArgument(argument).Flag())
	}
	appended.Action = a.action(command)

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

func (a *Application) addArguments(booter Booter, get get.Arguments) error {
	return a.container.Get(func(shell *core.Shell) {
		for _, argument := range get.Arguments {
			shell.Flags = append(shell.Flags, core.NewArgument(argument).Flag())
		}
		for _, argument := range core.NewTyper(booter).Arguments() {
			shell.Flags = append(shell.Flags, core.NewArgument(argument).Flag())
		}
	}).Build().Inject()
}

func (a *Application) addCommands(shell *core.Shell, booter Booter, get get.Commands) {
	for _, cmd := range get.Commands {
		shell.Commands = append(shell.Commands, a.convertCommand(cmd))
	}
	for _, cmd := range core.NewTyper(booter).Commands() {
		shell.Commands = append(shell.Commands, a.convertCommand(cmd))
	}
}

func (a *Application) setupDefaultAction(shell *core.Shell, booter Booter) { // 没有任何命令时执行动作
	shell.Action = func(ctx *cli.Context) (err error) {
		if converted, ok := booter.(checker.Run); ok {
			err = converted.Run(runtime.NewContext(ctx))
		}

		return
	}
}

func (a *Application) boot() (err error) {
	// 优雅退出
	go a.graceful(&err)

	canceled, cancel := context.WithTimeout(context.Background(), a.params.Timeout.Startup)
	defer cancel()

	if aae := a.container.Get(a.addArguments).Build().Inject(); nil != aae { // 增加参数
		err = aae
	} else if ace := a.container.Get(a.addCommands).Build().Inject(); nil != ace { // 增加命令
		err = ace
	} else if bpe := a.params.Banner.Print(); nil != bpe { // 打印标志信息
		err = bpe
	} else if ie := a.container.Get(a.initialize(canceled)).Build().Inject(); nil != ie { // 执行初始化方法
		err = ie
	} else if sle := a.container.Get(a.setLogger).Build().Inject(); nil != sle { // 为执行壳设置日志器
		err = sle
	} else if ple := a.putLogger(); nil != ple {
		err = ple
	} else if se := a.container.Get(a.startup(canceled)).Build().Inject(); nil != se { // 执行初始化方法
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

func (a *Application) startup(ctx context.Context) func(*core.Shell, Booter) error {
	return func(shell *core.Shell, booter Booter) (err error) {
		if bbe := a.beforeBooter(ctx, booter); nil != bbe { // 生命周期方法
			err = bbe
		} else if se := booter.Boot(ctx); nil != se { // 启动应用
			err = se
		} else if cbe := a.before(ctx); nil != cbe { // 执行命令生命周期方法
			err = cbe
		} else if ple := a.putLogger(); nil != ple { // 日志
			err = ple
		} else if re := shell.Run(os.Args); nil != re { // 执行整个程序
			err = re
		} else if cae := a.after(ctx); nil != cae { // 执行命令生命周期方法
			err = cae
		} else if abe := a.afterBooter(ctx, booter); nil != abe { // 执行命令生命周期方法
			err = abe
		}

		return
	}
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
	err = a.container.Get(a.setupShell).Build().Inject()

	return
}

func (a *Application) addDependency(constructor runtime.Constructor) (err error) {
	if ve := a.verify(constructor); nil != ve {
		err = ve
	} else if pce := a.container.Put(constructor).Invalidate().Build().Inject(); nil != pce {
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
	_ = a.container.Get(func(info *command.Info) error {
		return info.Run(runtime.NewContext(ctx))
	}).Build().Inject()
}

func (a *Application) putSelf() *Application {
	return a
}

func (a *Application) putLogger() (err error) {
	if ge := a.container.Get(a.getLogger).Build().Inject(); nil != ge {
		err = a.container.Put(a.supplyLogger).Build().Inject() // !当出错或未成功设置时，重置日志器
	}

	return
}

func (a *Application) getLogger(get get.Logger) {
	if nil != get.Logger { // !只有在确实有外部日志器的情况下才允许覆盖
		a.logger = get.Logger
	}
}

func (a *Application) supplyLogger() log.Logger {
	a.logger = log.New().Apply()
	if level, ok := os.LookupEnv(constant.EnvironmentLoggingLevel); ok {
		a.logger.Enable(log.ParseLevel(level))
	}

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
	} else if 1 != typ.NumOut() || reflect.TypeOf((*Booter)(nil)).Elem() != typ.Out(constant.IndexFirst) {
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

func (a *Application) beforeBooter(ctx context.Context, stater Booter) (err error) {
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

func (a *Application) afterBooter(ctx context.Context, stater Booter) (err error) {
	if converted, ok := stater.(application.After); ok {
		err = converted.After(ctx)
	}

	return
}
