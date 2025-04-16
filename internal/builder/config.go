package builder

import (
	"github.com/heluon/boot/internal/internal/builder"
	"github.com/heluon/boot/internal/internal/param"
)

// Config 配置构建器，编写非运行时的配置项
// 插件配置加载时也可以覆盖配置项
type Config struct {
	*builder.Config[Config]

	params      *param.Config
	application *Application
}

func newConfig(application *Application) (config *Config) {
	params := param.NewConfig()
	config = new(Config)
	config.Config = builder.NewConfig(params, config)
	config.params = params
	config.application = application

	return
}

func (c *Config) Validate() (config *Config) {
	c.params.Validate = true
	config = c

	return
}

func (c *Config) Invalidate() (config *Config) {
	c.params.Validate = false
	config = c

	return
}

func (c *Config) Required() (config *Config) {
	c.params.Nullable = false
	config = c

	return
}

func (c *Config) Nullable() (config *Config) {
	c.params.Nullable = true
	config = c

	return
}

func (c *Config) Build() (application *Application) {
	application = c.application
	if !c.params.Nullable {
		application.params.Config.Nullable = c.params.Nullable
	}
	if !c.params.Default {
		application.params.Config.Default = c.params.Default
	}
	if 0 != len(c.params.Environments) {
		application.params.Config.Environments = append(application.params.Config.Environments, c.params.Environments...)
	}
	if 0 != len(c.params.Loaders) {
		application.params.Config.Loaders = append(application.params.Config.Loaders, c.params.Loaders...)
	}
	for getter := range c.params.Getters {
		application.params.Config.Getters[getter] = c.params.Getters[getter]
	}

	return
}
