package param

import (
	"path/filepath"

	"github.com/goexl/env"
	"github.com/pangum/pangu/internal/callback/getter"
	"github.com/pangum/pangu/internal/config"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/app"
	"github.com/pangum/pangu/internal/internal/loader"
)

type Config struct {
	config.Loader
	config.Watcher

	Paths      []string
	Extensions []string

	Default  bool
	Validate bool
	Nullable bool

	Tag               *Tag
	EnvironmentGetter getter.Environment
	Environments      app.Environments
	Set               bool
}

func newConfig() (config *Config) {
	config = new(Config)
	config.Paths = []string{
		constant.ApplicationName,
		filepath.Join(constant.ConfigDir, constant.ApplicationName),
		filepath.Join(constant.ConfigConfDir, constant.ApplicationName),
		filepath.Join(constant.ConfigConfigurationDir, constant.ApplicationName),
	}
	config.Extensions = []string{
		constant.ExtensionYml,
		constant.ExtensionYaml,
		constant.ExtensionToml,
		constant.ExtensionJson,
		constant.ExtensionXml,
	}
	config.Default = true
	config.Validate = true
	config.Nullable = true

	config.Tag = NewTag()
	config.EnvironmentGetter = env.Get
	config.Environments = make(app.Environments, 0)
	config.Loader = loader.NewConfig(config.EnvironmentGetter, config.Nullable)

	return
}
