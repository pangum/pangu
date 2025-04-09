package param

import (
	"github.com/goexl/env"
	"github.com/pangum/config"
	"github.com/pangum/pangu/internal/callback"
	"github.com/pangum/pangu/internal/internal/kernel"
)

type Config struct {
	// 是否允许设置默认值
	Default bool
	// 是否要验证配置数据
	Validate bool
	// 是否可以没有配置文件
	Nullable bool
	// 是否可刷新配置
	Refreshable bool

	// 标签
	Tag *Tag
	// 环境变量获取器
	Getter callback.Environment
	// 环境亦是
	Environments kernel.Environments

	Loaders  []config.Loader
	Changers []config.Changer
}

func NewConfig(loaders ...config.Loader) *Config {
	return &Config{
		Default:     true,
		Validate:    true,
		Nullable:    true,
		Refreshable: true,

		Tag:          NewTag(),
		Getter:       env.Get,
		Environments: make(kernel.Environments, 0),

		Loaders:  loaders,
		Changers: make([]config.Changer, 0),
	}
}
