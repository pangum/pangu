package param

import (
	"path/filepath"

	"github.com/goexl/env"
	"github.com/goexl/mengpo"
	"github.com/goexl/xiren"
	"github.com/pangum/pangu/internal/callback/getter"
	"github.com/pangum/pangu/internal/config"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal"
)

type Config struct {
	Paths      []string
	Extensions []string

	Default  bool
	Validate bool
	Nullable bool

	Tag               *config.Tag
	EnvironmentGetter getter.Environment
	Environments      internal.Environments
	Watcher           config.Watcher
	Loader            config.Loader
}

func NewConfig() *Config {
	return &Config{
		Paths: []string{
			constant.ApplicationName,
			filepath.Join(constant.ConfigDir, constant.ApplicationName),
			filepath.Join(constant.ConfigConfDir, constant.ApplicationName),
			filepath.Join(constant.ConfigConfigurationDir, constant.ApplicationName),
		},
		Extensions: []string{
			constant.ExtensionYml,
			constant.ExtensionYaml,
			constant.ExtensionToml,
			constant.ExtensionJson,
			constant.ExtensionXml,
		},

		Default:  true,
		Validate: true,
		Nullable: true,

		Tag: &config.Tag{
			Default: constant.DefaultTag,
		},

		EnvironmentGetter: env.Get,
		Environments:      make(internal.Environments, 0),
	}
}

func (co *Config) Load(path string, config any) (err error) {
	if err = co.Loader.Load(path, config); nil != err {
		return
	}

	// 处理默认值，此处逻辑不能往前，原因
	// 如果对象里面包含指针，那么只能在包含指针的结构体被解析后才能去设置默认值，不然指针将被会设置成nil
	if co.Default {
		err = mengpo.New().Tag(co.Tag.Default).Build().Set(config)
	}
	if nil != err {
		return
	}

	// 数据验证
	if co.Validate {
		err = xiren.Struct(config)
	}

	return
}
