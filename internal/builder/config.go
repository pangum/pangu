package builder

import (
	"github.com/pangum/pangu/internal/internal/app"
	"github.com/pangum/pangu/internal/internal/builder"
	"github.com/pangum/pangu/internal/param"
)

type Config struct {
	*builder.Config

	params      *param.Config
	application *Application
}

func NewConfig(application *Application) (config *Config) {
	params := application.params.Config
	config = new(Config)
	config.Config = builder.NewConfig(params)
	config.params = params
	config.application = application

	return
}

func (c *Config) Path(path string) (config *Config) {
	c.params.Paths = append(c.params.Paths, path)
	config = c

	return
}

func (c *Config) Extension(extension string) (config *Config) {
	c.params.Extensions = append(c.params.Extensions, extension)
	config = c

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

func (c *Config) Environment(key string, value string) (config *Config) {
	c.params.Environments = append(c.params.Environments, app.NewEnvironment(key, value))
	config = c

	return
}

func (c *Config) Build() (application *Application) {
	c.params.Set = true
	application.params.Config = c.params
	application = c.application

	return
}
