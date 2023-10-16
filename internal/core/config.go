package core

import (
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/builder"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/urfave/cli/v2"
)

type Config struct {
	*builder.Config

	path     string
	original *param.Config
	params   *param.Config
}

func NewConfig(params *param.Config) (config *Config) {
	original := new(param.Config)
	*original = *params

	config = new(Config)
	config.Config = builder.NewConfig(original)
	config.original = params
	config.params = original

	return
}

func (c *Config) Load(config runtime.Pointer) (err error) {
	if path, fpe := c.filepath(); nil != fpe {
		err = fpe
	} else if fe := c.params.Fill(path, config); nil != fe { // 加载数据
		err = fe
	} else if nil != c.params.Watcher { // 配置文件监控
		// TODO err = c.Watch(config, c.params.Watcher)
	} else {
		c.path = path
	}

	return
}

func (c *Config) filepath() (path string, err error) {
	gfxOptions := gfx.NewExistsOptions(
		gfx.Paths(c.params.Paths...),
		gfx.Extensions(c.params.Extensions...),
	)
	// 如果配置了应用名称，可以使用应用名称的配置文件
	if constant.ApplicationDefaultName != runtime.Name {
		gfxOptions = append(gfxOptions, gfx.Paths(
			runtime.Name,
			filepath.Join(constant.ConfigDir, runtime.Name),
			filepath.Join(constant.ConfigConfDir, runtime.Name),
			filepath.Join(constant.ConfigConfigurationDir, runtime.Name),
		))
	}

	if final, exists := gfx.Exists(c.path, gfxOptions...); exists {
		path = final
	} else { // 如果找不到配置文件，则所用默认的配置文件
		path = c.path
	}

	return
}

func (c *Config) bind(shell *runtime.Shell, shadow *runtime.Shadow) {
	flag := new(cli.StringFlag)
	flag.Name = constant.ConfigName
	flag.Aliases = []string{
		constant.ConfigAliasC,
		constant.ConfigAliasConf,
		constant.ConfigAliasConfiguration,
	}
	flag.Value = constant.ConfigDefaultFilepath
	flag.Usage = "指定配置文件路径"
	flag.Destination = &c.path

	shell.Flags = append(shell.Flags, flag)
	shadow.Flags = append(shadow.Flags, flag)
}
