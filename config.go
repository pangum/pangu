package pangu

import (
	"path/filepath"

	"github.com/goexl/gfx"
	"github.com/urfave/cli/v2"
)

const (
	configDir        = `config`
	confDir          = `conf`
	configurationDir = `configuration`
)

// Config 配置处理器
type Config struct {
	// 路径
	path string
	// 选项
	options *configOptions
}

func newConfig(options *configOptions) *Config {
	return &Config{
		options: options,
	}
}

func (c *Config) Load(config any, opts ...configOption) (err error) {
	for _, opt := range opts {
		opt.applyConfig(c.options)
	}

	// 加载配置文件
	err = c.loadConfig(config)

	return
}

func (c *Config) Watch(config any, watcher configWatcher) (err error) {
	return gfx.Watch(c.path, newConfigFileWatcher(config, c.path, watcher, c.options))
}

func (c *Config) loadConfig(config any) (err error) {
	if c.path, err = c.configFilepath(); nil != err {
		return
	}

	// 加载数据
	if err = c.options.Load(c.path, config); nil != err {
		return
	}

	// 配置文件监控
	if nil != c.options.watcher {
		err = c.Watch(config, c.options.watcher)
	}

	return
}

func (c *Config) configFilepath() (path string, err error) {
	gfxOptions := gfx.NewExistsOptions(
		gfx.Paths(c.options.paths...),
		gfx.Extensions(c.options.extensions...),
	)
	// 如果配置了应用名称，可以使用应用名称的配置文件
	if defaultName != Name {
		gfxOptions = append(gfxOptions, gfx.Paths(
			Name,
			filepath.Join(configDir, Name),
			filepath.Join(confDir, Name),
			filepath.Join(configurationDir, Name),
		))
	}

	if final, exists := gfx.Exists(c.path, gfxOptions...); exists {
		path = final
	} else { // 如果找不到配置文件，则所用默认的配置文件
		path = c.path
	}

	return
}

func (c *Config) bind(shell *cli.App, shadow *cli.App) {
	configFlag := &cli.StringFlag{
		Name:        `config`,
		Aliases:     []string{`c`, `conf`, `configuration`},
		Value:       `./conf/application.yaml`,
		Usage:       `指定配置文件路径`,
		Destination: &c.path,
	}
	shell.Flags = append(shell.Flags, configFlag)
	shadow.Flags = append(shadow.Flags, configFlag)
}
