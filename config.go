package pangu

import (
	`encoding/json`
	`encoding/xml`

	`github.com/mcuadros/go-defaults`
	`github.com/pelletier/go-toml`
	`github.com/storezhang/validatorx`
	`gopkg.in/yaml.v3`
)

// Config 描述全局原始配置参数
type Config struct {
	// 原始数据
	data []byte
	// 格式
	format string
	// 全局选项
	options *options
}

func (c *Config) Struct(config interface{}, opts ...option) (err error) {
	for _, opt := range opts {
		opt.apply(c.options)
	}

	// 处理默认值
	if c.options.isDefault {
		defaults.SetDefaults(config)
	}

	switch c.format {
	case "yml":
		fallthrough
	case "yaml":
		err = yaml.Unmarshal(c.data, config)
	case "json":
		err = json.Unmarshal(c.data, config)
	case "toml":
		err = toml.Unmarshal(c.data, config)
	case "xml":
		err = xml.Unmarshal(c.data, config)
	default:
		err = yaml.Unmarshal(c.data, config)
	}
	if nil != err {
		return
	}

	// 验证数据
	if c.options.isValidate {
		err = validatorx.Struct(config)
	}

	return
}
