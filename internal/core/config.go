package core

import (
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/urfave/cli/v2"
)

type Config struct {
	path  string
	param *param.Config
}

func NewConfig(options *param.Config) *Config {
	return &Config{
		param: options,
	}
}

func (c *Config) Load(config any) (err error) {
	// 加载配置文件
	err = c.load(config)

	return
}

/* TODO func (c *Config) Watch(config any, watcher config.Watcher) (err error) {
	return gfx.Watch(c.path, pangu.newConfigFileWatcher(config, c.path, watcher, c.param))
}*/

func (c *Config) load(config any) (err error) {
	if c.path, err = c.configFilepath(); nil != err {
		return
	}

	// 加载数据
	if err = c.param.Load(c.path, config); nil != err {
		return
	}

	// 配置文件监控
	if nil != c.param.Watcher {
		// TODO err = c.Watch(config, c.param.Watcher)
	}

	return
}

func (c *Config) configFilepath() (path string, err error) {
	gfxOptions := gfx.NewExistsOptions(
		gfx.Paths(c.param.Paths...),
		gfx.Extensions(c.param.Extensions...),
	)
	// 如果配置了应用名称，可以使用应用名称的配置文件
	if constant.DefaultName != runtime.Name {
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
	flag.Name = "config"
	flag.Aliases = []string{"c", "conf", "configuration"}
	flag.Value = "./conf/application.yaml"
	flag.Usage = "指定配置文件路径"
	flag.Destination = &c.path

	shell.Flags = append(shell.Flags, flag)
	shadow.Flags = append(shadow.Flags, flag)
}
