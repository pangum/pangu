package conf

import (
	`github.com/storezhang/pangu`
)

// Config 总的配置入口
// 支持以下几种配置
// Yaml
// Toml
// Json
// XML
type Config struct {
	// 本应用程序配置
	Example Example `validate:"required"`
}

// 暴露配置，可以是私有也可以仅有方法，建议私有，尽量隐藏现实
func config(application *pangu.Application) (config Config, err error) {
	err = application.LoadConfig(&config)

	return
}

// 暴露配置（快捷方法），也可以选择不暴露，那么在使用的时候就以config.Example这种形式去使用
func example(config Config) Example {
	return config.Example
}
