package config

import (
	"sync"

	"github.com/goexl/exception"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/log"
	"github.com/goexl/mengpo"
	"github.com/goexl/xiren"
	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/config/internal/core"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/urfave/cli/v2"
)

type Getter struct {
	path   gox.Pointer[string]
	params *param.Config

	logger *log.Logger
	once   *sync.Once

	environment *core.Environment
	loader      *core.Loader
	watcher     *core.Watcher
}

func newGetter(params *param.Config, logger *log.Logger) (getter *Getter) {
	getter = new(Getter)
	getter.path = gox.NewPointer("")
	getter.params = params

	getter.logger = logger
	getter.once = new(sync.Once)

	getter.environment = core.NewEnvironment()
	getter.loader = core.NewLoader(getter.path, params, logger)
	getter.watcher = core.NewWatch(getter.loader)

	return
}

func (g *Getter) Get(target runtime.Pointer) (err error) {
	if fpe := g.filepath(); nil != fpe { // 获取最终的配置文件路径
		err = fpe
	} else if fe := g.fill(target); nil != fe { // 加载数据
		err = fe
	}

	return
}

func (g *Getter) fill(target runtime.Pointer) (err error) {
	if le := g.loader.Load(target); nil != le { // 从路径中加载数据
		err = le
	}

	if nil == err && g.params.Default { // 处理默认值
		// !此处逻辑不能往前，原因是如果对象里面包含指针，那么只能在包含指针的结构体被解析后才能去设置默认值，不然指针将被会设置成空值
		err = mengpo.New().Tag(g.params.Tag.Default).Build().Set(target)
	}

	// 从环境变量中加载配置
	if nil == err {
		err = g.environment.Process(target)
	}

	if nil == err && g.params.Validate { // 数据验证
		err = xiren.Struct(target)
	}

	return
}

func (g *Getter) filepath() (err error) {
	if nil == g.path || "" == *g.path {
		err = g.detectFilepath()
	}

	return
}

func (g *Getter) detectFilepath() (err error) {
	exists := gfx.Exists().Filename(constant.ApplicationName)
	if constant.EnvironmentNotSet != internal.Name { // 如果配置了应用名称，可以使用应用名称的配置文件
		exists.Filename(internal.Name)
	}
	// 配置所有可能的配置目录
	exists.Directory("config")
	exists.Directory("conf")
	exists.Directory("configuration")

	if final, checked := exists.Build().Check(); checked {
		*g.path = final // !一定要使用指针修改原来的值，而不是传新的指针
	} else if !g.params.Nullable {
		err = exception.New().Message("配置文件不存在").Build()
	}

	return
}

func (g *Getter) bind(shell *runtime.Shell, shadow *runtime.Shadow) {
	config := new(cli.StringFlag)
	config.Name = constant.ConfigName
	config.Aliases = []string{
		constant.ConfigAliasC,
		constant.ConfigAliasConf,
		constant.ConfigAliasConfiguration,
	}
	config.Usage = "指定配置文件路径"
	config.Destination = g.path

	shell.Flags = append(shell.Flags, config)
	shadow.Flags = append(shadow.Flags, config)
}
