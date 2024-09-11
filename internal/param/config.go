package param

import (
	"github.com/goexl/env"
	"github.com/pangum/config"
	"github.com/pangum/pangu/internal/callback"
	"github.com/pangum/pangu/internal/internal/app"
	"github.com/pangum/pangu/internal/internal/loader"
)

type Config struct {
	Default  bool
	Validate bool
	Nullable bool

	Tag               *Tag
	EnvironmentGetter callback.Environment
	Environments      app.Environments

	Loaders []config.Loader
}

func newConfig() *Config {
	return &Config{
		Default:           true,
		Validate:          true,
		Nullable:          true,
		Tag:               NewTag(),
		EnvironmentGetter: env.Get,
		Environments:      make(app.Environments, 0),

		Loaders: []config.Loader{
			loader.NewJson(),
			loader.NewXml(),
		},
	}
}
