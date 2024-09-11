package builder

import (
	"github.com/pangum/config"
	"github.com/pangum/pangu/internal/internal/app"
	"github.com/pangum/pangu/internal/param"
)

type Config struct {
	param *param.Config
}

func NewConfig(param *param.Config) *Config {
	return &Config{
		param: param,
	}
}

func (c *Config) Nullable() (config *Config) {
	c.param.Nullable = true
	config = c

	return
}

func (c *Config) Required() (config *Config) {
	c.param.Nullable = false
	config = c

	return
}

func (c *Config) Default() (config *Config) {
	c.param.Default = true
	config = c

	return
}

func (c *Config) Environment(key string, value string) (config *Config) {
	c.param.Environments = append(c.param.Environments, app.NewEnvironment(key, value))
	config = c

	return
}

func (c *Config) Loader(loader config.Loader, loaders ...config.Loader) (config *Config) {
	c.param.Loaders = append(c.param.Loaders, loader)
	c.param.Loaders = append(c.param.Loaders, loaders...)
	config = c

	return
}
