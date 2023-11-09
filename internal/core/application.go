package core

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"

	"github.com/goexl/exc"
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
	"github.com/pangum/pangu/internal/message"
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
	application.config = config.NewSetup(params.Config)
	application.stoppers = make([]app.Stopper, 0)
	application.lifecycles = make([]app.Lifecycle, 0)
	if err := application.addCore(); nil != err {
		panic(err)
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
	if ese := a.params.Environments.Set(); nil != ese { // 添加环境变量
		err = ese
	} else if ape := a.addDependency(constructor); nil != ape { // 添加内置的依赖
		err = ape
	} else if bde := dependency.Get(a.bind).Build().Build().Inject(); nil != bde { // 绑定参数和命令到内部变量或者命令上
		err = bde
	} else if cae := a.createApp(); nil != cae { // 创建应用
		err = cae
	} else if ace := dependency.Get(a.addCommands).Build().Build().Inject(); nil != ace { // 增加内置的命令及参数
		err = ace
	} else if bpe := a.params.Banner.Print(); nil != bpe { // 输出标志信息
		err = bpe
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
			err = exc.NewField("不支持组件类型", field.New("type", reflect.TypeOf(component).String()))
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

func (a *Application) bind(shell *runtime.Shell) (err error) {
	// 接收配置文件路径参数
	a.config.Bind(shell, a.shadow)
	// 影子执行，只有这样才能正确的使用配置文件的路径参数
	err = a.shadow.Run(a.args())

	return
}

func (a *Application) boot(bootstrap Bootstrap) (err error) {
	// 优雅退出
	go a.graceful(&err)

	canceled, cancel := context.WithTimeout(context.Background(), a.params.Timeout.Boot)
	defer cancel()

	if bse := bootstrap.Startup(a); nil != bse { // 加载用户启动器并做好配置
		err = bse
	} else if bbe := bootstrap.Before(canceled); nil != bbe { // 执行生命周期方法
		err = bbe
	} else if cbe := a.before(canceled); nil != cbe { // 执行命令生命周期方法
		err = cbe
	} else if re := a.Dependency().Get(a.run).Build().Build().Inject(); nil != re { // 执行整个程序
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

func (a *Application) addCore() error {
	return a.Dependency().Put(
		a.putConfig,      // 注入配置
		a.putSelf,        // 注入自身
		runtime.NewShell, // 注入运行壳
	).Build().Invalidate().Build().Inject()
}

func (a *Application) setupApp(shell *runtime.Shell) {
	shell.Name = internal.Name
	shell.Description = a.params.Description
	shell.Usage = a.params.Usage
	shell.Copyright = a.params.Copyright
	shell.Metadata = a.params.Metadata
	shell.Authors = a.params.Authors.Cli()
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
		a.logger = log.New().Apply()
		err = dependency.Put(a.supplyLogger).Build().Build().Inject()
	}

	return
}

func (a *Application) getLogger(logger log.Logger) {
	a.logger = logger
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
		err = exc.NewField(message.ConstructorMustFunc, field.New("bootstrap", typ.String()))
	} else if 0 == typ.NumIn() { // 构造方法必须有依赖项
		constructorName := runtime.FuncForPC(reflect.ValueOf(bootstrap).Pointer()).Name()
		err = exc.NewField(message.BootstrapMustHasDependencies, field.New("bootstrap", constructorName))
	} else if 1 != typ.NumOut() || reflect.TypeOf((*Bootstrap)(nil)).Elem() != typ.Out(constant.IndexFirst) {
		// 只能返回一个类型为Bootstrap返回值
		err = exc.NewMessage(message.BootstrapMustReturnBootstrap)
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

func (a *Application) before(ctx context.Context) (err error) {
	a.logger.Info(
		"准备启动应用",
		field.New("name", internal.Name),
		field.New("version", internal.Version),
		field.New("build", internal.Build),
		field.New("datetime", internal.Timestamp),
		field.New("revision", internal.Revision),
		field.New("branch", internal.Branch),
		field.New("runtime", internal.Runtime),
	)
	for _, after := range a.lifecycles {
		err = after.Before(ctx)
		if nil != err {
			break
		}
	}

	return
}

func (a *Application) after(ctx context.Context) (err error) {
	for _, after := range a.lifecycles {
		err = after.After(ctx)
		if nil != err {
			break
		}
	}

	return
}
