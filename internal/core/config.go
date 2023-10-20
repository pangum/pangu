package core

import (
	"github.com/pangum/pangu/internal/internal/builder"
	"github.com/pangum/pangu/internal/internal/config"
	"github.com/pangum/pangu/internal/param"
)

type Config struct {
	*builder.Config

	setup    *config.Setup
	original *param.Config
	params   *param.Config
	override bool
}

func NewConfig(setup *config.Setup, params *param.Config) (config *Config) {
	original := new(param.Config)
	*original = *params

	config = new(Config)
	config.setup = setup
	config.Config = builder.NewConfig(original)
	config.params = params

	return
}

func (c *Config) Override() (config *Config) {
	c.override = true
	config = c

	return
}

func (c *Config) Build() (getter *config.Getter) {
	getter = c.setup.Copy(c.params)
	if c.override {
		*c.params = *c.original
	}

	return
}
