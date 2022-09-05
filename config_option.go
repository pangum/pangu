package pangu

import (
	"path/filepath"

	"github.com/goexl/env"
	"github.com/goexl/mengpo"
	"github.com/goexl/xiren"
)

type (
	configOption interface {
		applyConfig(options *configOptions)
	}

	configOptions struct {
		paths      []string
		extensions []string

		defaults  bool
		validates bool
		nullable  bool

		tag               tag
		environmentGetter environmentGetter
		environments      []*environment
		watcher           configWatcher
		loader            configLoader
	}
)

func defaultConfigOptions() (_options *configOptions) {
	_options = &configOptions{
		paths: []string{
			applicationName,
			filepath.Join(configDir, applicationName),
			filepath.Join(confDir, applicationName),
			filepath.Join(configurationDir, applicationName),
		},
		extensions: []string{
			ymlExt,
			yamlExt,
			tomlExt,
			jsonExt,
			xmlExt,
		},

		defaults:  true,
		validates: true,
		nullable:  true,

		tag: tag{
			defaults: defaultsTag,
		},

		environmentGetter: env.Get,
		environments:      make([]*environment, 0),
	}
	_options.loader = newDefaultConfigLoader(_options)

	return
}

func (co *configOptions) Load(path string, config any) (err error) {
	if err = co.loader.Load(path, config); nil != err {
		return
	}

	// 处理默认值，此处逻辑不能往前，原因
	// 如果对象里面包含指针，那么只能在包含指针的结构体被解析后才能去设置默认值，不然指针将被会设置成nil
	if co.defaults {
		err = mengpo.Set(config, mengpo.Tag(co.tag.defaults))
	}
	if nil != err {
		return
	}

	// 数据验证
	if co.validates {
		err = xiren.Struct(config)
	}

	return
}
