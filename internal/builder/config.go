package builder

import (
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

func (c *Config) Build() (application *Application) {
	c.params.Set = true
	application.params.Config = c.params
	application = c.application

	return
}
