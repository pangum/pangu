package pangu

import (
	`encoding/json`
	`encoding/xml`
	`errors`
	`flag`
	`fmt`
	`io/ioutil`
	`path/filepath`
	`strings`
	`sync`

	`github.com/drone/envsubst`
	`github.com/pelletier/go-toml`
	`github.com/storezhang/gox`
	`github.com/storezhang/mengpo`
	`github.com/storezhang/validatorx`
	`gopkg.in/yaml.v3`
)

// Config 描述全局原始配置参数
type Config struct {
	path *string
	// 原始数据
	data []byte
	// 格式
	format string
	// 全局程序
	application *Application
	// 单例模式
	once sync.Once
}

func (c *Config) Load(config interface{}, opts ...option) (err error) {
	for _, opt := range opts {
		opt.apply(c.application.options)
	}

	// 参数不允许重复定义，只能执行一次
	c.once.Do(func() {
		c.path = flag.String(`conf`, `./conf/application.yaml`, `指定配置文件路径`)
		flag.StringVar(c.path, `c`, *c.path, `指定配置文件路径`)
		flag.Parse()
	})
	err = c.loadConfig(config)

	return
}

func (c *Config) loadConfig(config interface{}) (err error) {
	var finalPath string
	if finalPath, err = c.findConfigFilepath(*c.path); nil != err {
		return
	}

	// 可以处理后续动态加载
	if `` == c.format {
		c.format = strings.ToLower(filepath.Ext(finalPath))
	}
	if 0 == len(c.data) {
		if c.data, err = ioutil.ReadFile(finalPath); nil != err {
			return
		}
	}

	// 处理环境变量
	if err = c.eval(); nil != err {
		return
	}

	switch c.format {
	case `.yml`:
		fallthrough
	case `.yaml`:
		err = yaml.Unmarshal(c.data, config)
	case `.json`:
		err = json.Unmarshal(c.data, config)
	case `.toml`:
		err = toml.Unmarshal(c.data, config)
	case `.xml`:
		err = xml.Unmarshal(c.data, config)
	default:
		err = yaml.Unmarshal(c.data, config)
	}
	if nil != err {
		return
	}

	// 处理默认值，此处逻辑不能往前，原因
	// 如果对象里面包含指针，那么只能在包含指针的结构体被解析后才能去设置默认值，不然指针将被会设置成nil
	if c.application.options._default {
		if err = mengpo.Set(config, mengpo.Tag(c.application.options.tag._default)); nil != err {
			return
		}
	}

	// 验证数据
	if c.application.options.validate {
		err = validatorx.Struct(config)
	}

	return
}

// 环境变量配置处理
// 后续增加表达式处理
func (c *Config) eval() (err error) {
	var raw string
	if raw, err = envsubst.EvalEnv(string(c.data)); nil != err {
		return
	}
	c.data = []byte(raw)

	return
}

func (c *Config) findConfigFilepath(conf string) (path string, err error) {
	path = conf
	if `` != path && gox.IsFileExist(path) {
		return
	}

	var notExists bool
	if path, notExists = c.findConfigFilepathWithExt(`./application`); !notExists {
		return
	}
	if path, notExists = c.findConfigFilepathWithExt(`./conf/application`); !notExists {
		return
	}
	if path, notExists = c.findConfigFilepathWithExt(`./name`); !notExists {
		return
	}
	if path, notExists = c.findConfigFilepathWithExt(`./conf/name`); !notExists {
		return
	}
	err = errors.New(`找不到配置文件`)

	return
}

// 之所以命名为notExists，是为了少对notExists赋值
func (c *Config) findConfigFilepathWithExt(filename string) (path string, notExists bool) {
	path = fmt.Sprintf(`%s.%s`, filename, `yaml`)
	if gox.IsFileExist(path) {
		return
	}

	path = fmt.Sprintf(`%s.%s`, filename, `yml`)
	if gox.IsFileExist(path) {
		return
	}

	path = fmt.Sprintf(`%s.%s`, filename, `toml`)
	if gox.IsFileExist(path) {
		return
	}

	path = fmt.Sprintf(`%s.%s`, filename, `json`)
	if gox.IsFileExist(path) {
		return
	}

	path = fmt.Sprintf(`%s.%s`, filename, `xml`)
	if gox.IsFileExist(path) {
		return
	}
	notExists = true

	return
}
