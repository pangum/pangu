package param

import (
	"github.com/goexl/env"
	"github.com/goexl/gox"
	"github.com/harluo/boot/internal/internal/callback"
	"github.com/harluo/boot/internal/internal/kernel"
	"github.com/harluo/boot/internal/internal/param/internal/getter"
	"github.com/harluo/config"
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
	// 配置文件列表
	Paths []string

	// 标签
	Tag *Tag
	// 环境变量获取器
	Getters map[callback.Getter]*gox.Empty
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
		Paths:       make([]string, 0), // 默认没有配置文件

		Tag: NewTag(),
		Getters: map[callback.Getter]*gox.Empty{
			getter.NewDefault(env.Get): new(gox.Empty),
		},
		Environments: make(kernel.Environments, 0),

		Loaders:  loaders,
		Changers: make([]config.Changer, 0),
	}
}

func (c *Config) Get(key string) (value string) {
	for _getter := range c.Getters {
		value = _getter.Get(key)
		if "" != value { // 及时回退
			break
		}
	}

	return
}
