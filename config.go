package pangu

import (
	`encoding/json`
	`encoding/xml`
	`errors`
	`flag`
	`io/ioutil`
	`path/filepath`
	`strings`
	`sync`

	`github.com/goexl/gfx`
	`github.com/goexl/mengpo`
	`github.com/goexl/xiren`

	`github.com/pelletier/go-toml`
	`gopkg.in/yaml.v3`
)

// Config 描述全局原始配置参数
type Config struct {
	path string
	// 原始数据
	data []byte
	// 选项
	options *configOptions
	// 单例模式
	once sync.Once
}

func (c *Config) Load(config interface{}, opts ...configOption) (err error) {
	for _, opt := range opts {
		opt.applyConfig(c.options)
	}

	// 参数不允许重复定义，只能执行一次
	c.once.Do(func() {
		c.path = *flag.String(configLongName, configDefault, configUsage)
		flag.StringVar(&c.path, configShortName, c.path, configUsage)
		flag.Parse()
	})
	err = c.loadConfig(config)

	return
}

func (c *Config) loadConfig(config interface{}) (err error) {
	if path, existErr := c.configFilepath(c.path); nil != err {
		err = existErr
	} else if c.loadable() {
		c.path = path
		c.data, err = ioutil.ReadFile(path)
	}
	if nil != err {
		return
	}

	switch strings.ToLower(filepath.Ext(c.path)) {
	case ymlExt:
		fallthrough
	case yamlExt:
		err = yaml.Unmarshal(c.data, config)
	case jsonExt:
		err = json.Unmarshal(c.data, config)
	case tomlExt:
		err = toml.Unmarshal(c.data, config)
	case xmlExt:
		err = xml.Unmarshal(c.data, config)
	default:
		err = yaml.Unmarshal(c.data, config)
	}
	if nil != err {
		return
	}

	// 处理默认值，此处逻辑不能往前，原因
	// 如果对象里面包含指针，那么只能在包含指针的结构体被解析后才能去设置默认值，不然指针将被会设置成nil
	if c.options.defaults {
		if err = mengpo.Set(config, mengpo.Tag(c.options.tag.defaults)); nil != err {
			return
		}
	}

	// 验证数据
	if c.options.validates {
		err = xiren.Struct(config)
	}

	return
}

func (c *Config) configFilepath(conf string) (path string, err error) {
	gfxOptions := gfx.NewExistsOptions(
		gfx.Path(c.options.paths[0], c.options.paths[1:]...),
		gfx.Ext(c.options.extensions[0], c.options.extensions[1:]...),
	)
	if final, exists := gfx.Exists(conf, gfxOptions...); exists {
		path = final
	} else {
		err = errors.New(`找不到配置文件`)
	}

	return
}

func (c *Config) loadable() bool {
	return `` == c.path || nil == c.data
}
