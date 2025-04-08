package param

import (
	"github.com/goexl/env"
	"github.com/pangum/config"
	"github.com/pangum/pangu/internal/callback"
	"github.com/pangum/pangu/internal/internal/core"
)

type Config struct {
	Default     bool
	Validate    bool
	Nullable    bool
	Refreshable bool

	Tag               *Tag
	EnvironmentGetter callback.Environment
	Environments      core.Environments

	Loaders  []config.Loader
	Changers []config.Changer
}

func NewConfig(loaders ...config.Loader) *Config {
	return &Config{
		Default:     true,
		Validate:    true,
		Nullable:    true,
		Refreshable: true,

		Tag:               NewTag(),
		EnvironmentGetter: env.Get,
		Environments:      make(core.Environments, 0),

		Loaders:  loaders,
		Changers: make([]config.Changer, 0),
	}
}
