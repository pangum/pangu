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
	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/command"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/builder"
	"github.com/pangum/pangu/internal/internal/config"
	"github.com/pangum/pangu/internal/internal/get"
	"github.com/pangum/pangu/internal/internal/message"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

var shadow *Application

type Application struct {
	config    *config.Setup
	params    *param.Application
	container *dig.Container
	// 影子启动器，用来处理额外的命令或者参数，因为正常的启动器无法完成此操作，原因是
	// 正常的启动器只提供一个Run方法来处理传数的参数，而此方法一旦执行，就意味着内部的命令也开始执行，而此时依赖关系还没有准备好
	shadow *runtime.Shadow
	// 存储所有可被停止的命令或者服务
	stoppers   []app.Stopper
	lifecycles []app.Lifecycle
	logger     log.Logger
}

func New(param *param.Application) (application *Application) {
	once.Do(create(param))
	application = shadow

	return
}

func create(params *param.Application) func() {
	return func() {
		shadow = new(Application)
		shadow.params = params
		shadow.container = dig.New()
		shadow.shadow = runtime.NewShadow()
		shadow.logger = shadow.getInternalLogger()
		shadow.config = config.NewSetup(params.Config, &shadow.logger)
		shadow.stoppers = make([]app.Stopper, 0)
		shadow.lifecycles = make([]app.Lifecycle, 0)

		// ! 这个操作必须在创建的时候就执行，因为在后续插件启动时会寻找下面的依赖，而在这个时候启动方法还没有执行
		shadow.Dependency().Puts(
			runtime.NewShell, // 注入运行壳
			shadow.putConfig, // 注入配置
			shadow.putSelf,   // 注入自身
		).Get(
			shadow.bind, // 绑定参数和命令到内部变量或者命令上
		).Build().Invalidate().Build().Apply()
	}
}

func (a *Application) Dependency() *builder.Dependency {
	return builder.NewDependency(a.container, a.params)
}

// Run 启动应用程序
func (a *Application) Run(constructor runtime.Constructor) {
	var err error
	defer a.finally(&err)

	dependency := a.Dependency()
	if bpe := a.params.Banner.Print(); nil != bpe { // 输出标志信息
		err = bpe
	} else if ese := a.params.Environments.Set(); nil != ese { // 添加环境变量
		err = ese
	} else if ape := a.addDependency(constructor); nil != ape { // 添加内置的依赖
		err = ape
	} else if cae := a.createApp(); nil != cae { // 创建应用
		err = cae
	} else if ace := dependency.Get(a.addCommands).Build().Build().Inject(); nil != ace { // 增加内置的命令及参数
		err = ace
	} else {
		err = dependency.Get(a.boot).Build().Build().Inject()
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
			typ := field.New("type", reflect.TypeOf(component).String())
			err = exception.New().Message(message.ComponentNotSupport).Field(typ).Build()
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

func (a *Application) addServe(serve app.Serve) error {
	return a.Dependency().Get(func(command *command.Serve) {
		command.Add(serve)
		a.stoppers = append(a.stoppers, serve)
	}).Build().Build().Inject()
}

func (a *Application) addCommand(command app.Command) error {
	return a.Dependency().Get(func(shell *runtime.Shell) {
		shell.Commands = append(shell.Commands, &cli.Command{
			Name:        command.Name(),
			Aliases:     command.Aliases(),
			Usage:       command.Usage(),
			Description: command.Description(),
			Subcommands: command.Subcommands().Cli(),
			Category:    command.Category(),
			Flags:       command.Arguments().Flags(),
			Hidden:      command.Hidden(),
			Action:      a.action(command),
		})
		a.stoppers = append(a.stoppers, command)
		a.lifecycles = append(a.lifecycles, command)
	}).Build().Build().Inject()
}

func (a *Application) addArg(argument app.Argument) error {
	return a.Dependency().Get(func(shell *runtime.Shell) {
		shell.Flags = append(shell.Flags, argument.Flag())
	}).Build().Build().Inject()
}

func (a *Application) boot(bootstrap Bootstrap) (err error) {
	// 优雅退出
	go a.graceful(&err)

	canceled, cancel := context.WithTimeout(context.Background(), a.params.Timeout.Boot)
	defer cancel()

	dependency := a.Dependency()
	if sle := dependency.Get(a.setLogger).Build().Build().Inject(); nil != sle { // 为执行壳设置日志器
		err = sle
	} else if bse := bootstrap.Startup(a); nil != bse { // 加载用户启动器并做好配置
		err = bse
	} else if bbe := bootstrap.Before(canceled); nil != bbe { // 执行生命周期方法
		err = bbe
	} else if cbe := a.before(canceled); nil != cbe { // 执行命令生命周期方法
		err = cbe
	} else if re := dependency.Get(a.run).Build().Build().Inject(); nil != re { // 执行整个程序
		err = re
	} else if cae := a.after(canceled); nil != cae { // 执行命令生命周期方法
		err = cae
	} else if bae := bootstrap.After(canceled); nil != bae { // 执行生命周期方法
		err = bae
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
	err = a.Dependency().Get(a.setupApp).Build().Build().Inject()

	return
}

func (a *Application) addCommands(get get.Command) (err error) {
	if se := a.addCommand(get.Serve); nil != se {
		err = se
	} else if ie := a.addCommand(get.Info); nil != ie {
		err = ie
	} else {
		err = a.addCommand(get.Version)
	}

	return
}

func (a *Application) addDependency(constructor runtime.Constructor) (err error) {
	dependency := a.Dependency().Invalidate()
	if ve := a.verify(constructor); nil != ve {
		err = ve
	} else if ple := a.putLogger(); nil != ple {
		err = ple
	} else if pse := dependency.Put(command.NewServe).Build().Build().Inject(); nil != pse { // 注入服务命令
		err = pse
	} else if pie := dependency.Put(command.NewInfo).Build().Build().Inject(); nil != pie { // 注入信息命令
		err = pie
	} else if pve := dependency.Put(command.NewVersion).Build().Build().Inject(); nil != pve { // 注入版本命令
		err = pve
	} else if pce := dependency.Put(constructor).Build().Build().Inject(); nil != pce {
		err = pce
	}

	return
}

func (a *Application) bind(shell *runtime.Shell) (err error) {
	// 接收配置文件路径参数
	a.config.Bind(shell, a.shadow)
	// 组装影子应用执行的参数
	originals := a.args()
	args := make([]string, 0, 3)
	args = append(args, originals[0])
	for index := 1; index < len(a.args()); index++ {
		argument := strings.ReplaceAll(originals[index], constant.StringStrike, constant.Empty)
		if a.isConfigArgument(argument) { // 只接收和配置相关的参数
			args = append(args, originals[index], originals[index+1])
		}
	}
	if 3 >= len(args) { // ! 检查是否是只有配置相关参数，没有任何其它命令，加入一个不存在的命令调用，防止没有传任何命令时系统给出用法提示
		args = append(args, constant.CommandNonexistent)
	}
	// 影子执行，只有这样才能正确的使用配置文件的路径参数
	err = a.shadow.RunContext(context.Background(), args)

	return
}

func (a *Application) setupApp(shell *runtime.Shell) {
	shell.Name = internal.Name
	shell.Description = a.params.Description
	shell.Usage = a.params.Usage
	shell.Copyright = a.params.Copyright
	shell.Metadata = a.params.Metadata
	shell.Authors = a.params.Authors.Cli()
	shell.AllowExtFlags = true
}

func (a *Application) versionPrinter(ctx *cli.Context) {
	_ = a.Dependency().Get(func(info *command.Info) error {
		return info.Run(runtime.NewContext(ctx))
	}).Build().Build().Inject()
}

func (a *Application) putConfig() *Config {
	return NewConfig(a.config, a.params.Config)
}

func (a *Application) putSelf() *Application {
	return a
}

func (a *Application) putLogger() (err error) {
	dependency := a.Dependency()
	if gle := dependency.Get(a.getLogger).Build().Build().Inject(); nil != gle {
		err = dependency.Put(a.supplyLogger).Build().Build().Inject()
	}

	return
}

func (a *Application) getLogger(logger log.Logger) {
	a.logger = logger
}

func (a *Application) getInternalLogger() (logger log.Logger) {
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
	if a.params.Verify {
		return
	}

	typ := reflect.TypeOf(bootstrap)
	if reflect.Func != typ.Kind() { // 构造方法必须是方法不能是其它类型
		err = exception.New().Message(message.ConstructorMustFunc).Field(field.New("bootstrap", typ.String())).Build()
	} else if 0 == typ.NumIn() { // 构造方法必须有依赖项
		name := runtime.FuncForPC(reflect.ValueOf(bootstrap).Pointer()).Name()
		err = exception.New().Message(message.BootstrapMustHasDependencies).Field(field.New("bootstrap", name)).Build()
	} else if 1 != typ.NumOut() || reflect.TypeOf((*Bootstrap)(nil)).Elem() != typ.Out(constant.IndexFirst) {
		// 只能返回一个类型为Bootstrap返回值
		err = exception.New().Message(message.BootstrapMustReturnBootstrap).Build()
	}

	return
}

func (a *Application) action(command app.Command) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return command.Run(runtime.NewContext(ctx))
	}
}

func (a *Application) graceful(err *error) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
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

	return
}

func (a *Application) setLogger(shell *runtime.Shell) {
	shell.Logger(a.logger)
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
	for _, lifecycle := range a.lifecycles {
		err = lifecycle.Before(ctx)
		if nil != err {
			break
		}
	}

	return
}

func (a *Application) after(ctx context.Context) (err error) {
	for _, lifecycle := range a.lifecycles {
		err = lifecycle.After(ctx)
		if nil != err {
			break
		}
	}

	return
}

func (a *Application) isConfigArgument(argument string) bool {
	return constant.ConfigName == argument || constant.ConfigC == argument || constant.ConfigConf == argument
}
