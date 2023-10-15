package builder

import (
	"github.com/pangum/pangu/internal/internal"
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

func (c *Config) Path(path string) (config *Config) {
	c.param.Paths = append(c.param.Paths, path)
	config = c

	return
}

func (c *Config) Extension(extension string) (config *Config) {
	c.param.Extensions = append(c.param.Extensions, extension)
	config = c

	return
}

func (c *Config) Default() (config *Config) {
	c.param.Default = true
	config = c

	return
}

func (c *Config) Environment(key string, value string) (config *Config) {
	c.param.Environments = append(c.param.Environments, internal.NewEnvironment(key, value))
	config = c

	return
}
