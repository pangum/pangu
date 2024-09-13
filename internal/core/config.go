package core

import (
	"github.com/goexl/gox"
	"github.com/pangum/pangu/internal/internal/builder"
	"github.com/pangum/pangu/internal/internal/config"
	"github.com/pangum/pangu/internal/param"
)

type Config struct {
	*builder.Config[Config]

	setup    *config.Setup
	original *param.Config
	params   *param.Config
	override bool

	_ gox.Pointerized
}

func NewConfig(setup *config.Setup, params *param.Config) (config *Config) {
	original := new(param.Config)
	*original = *params

	config = new(Config)
	config.setup = setup
	config.Config = builder.NewConfig(original, config)
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
