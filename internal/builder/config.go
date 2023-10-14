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

func (c *Config) Path(paths ...string) (config *Config) {
	c.param.Paths = append(c.param.Paths, paths...)
	config = c

	return
}

func (c *Config) Extension(extensions ...string) (config *Config) {
	c.param.Extensions = append(c.param.Extensions, extensions...)
	config = c

	return
}

func (c *Config) Default() (config *Config) {
	c.param.Default = true
	config = c

	return
}

func (c *Config) Environment(key string, value string) (config *Config) {
	author := new(internal.Environment)
	author.Key = key
	author.Value = value

	c.param.Environments = append(c.param.Environments, author)
	config = c

	return
}
