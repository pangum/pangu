package core

import (
	"fmt"
	"os"
	"reflect"
	"sync"

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

func (a *Application) addServe(serves ...app.Serve) error {
	return a.Dependency().Get(func(cmd *command.Serve) {
		for _, serve := range serves {
			cmd.Add(serve)
		}
	}).Build().Build().Inject()
}

func (a *Application) addCommand(commands ...app.Command) error {
	return a.Dependency().Get(func(shell *runtime.Shell) {
		shell.Commands = append(shell.Commands, app.Commands(commands).Cli()...)
	}).Build().Build().Inject()
}

func (a *Application) addArg(args ...app.Argument) error {
	return a.Dependency().Get(func(shell *runtime.Shell) {
		for _, arg := range args {
			_arg := arg
			shell.Flags = append(shell.Flags, _arg.Flag())
		}
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
	if bse := bootstrap.Startup(a); nil != bse { // 加载用户启动器并做好配置
		err = bse
	} else if rbe := bootstrap.Before(); nil != rbe { // 执行生命周期方法
		err = rbe
	} else if re := a.Dependency().Get(a.run).Build().Build().Inject(); nil != re { // 执行整个程序
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
	err = a.Dependency().Get(a.setupApp).Build().Build().Inject()

	return
}

func (a *Application) addCommands(get get.Command) error {
	return a.addCommand(get.Serve, get.Info, get.Version)
}

func (a *Application) addDependency(constructor runtime.Constructor) (err error) {
	dependency := a.Dependency().Invalidate()
	if ve := a.verify(constructor); nil != ve {
		err = ve
	} else if ple := dependency.Put(a.putLogger).Build().Build().Inject(); nil != ple {
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

func (a *Application) putLogger() (logger app.Logger) {
	err := a.Dependency().Get(func(external log.Logger) {
		logger = external
	}).Build().Build().Inject()
	if nil != err {
		logger = log.New().Apply()
	}

	return
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
