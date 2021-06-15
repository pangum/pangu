package pangu

import (
	`encoding/json`
	`encoding/xml`
	`errors`
	`flag`
	`fmt`
	`io/ioutil`
	`path/filepath`
	`reflect`
	`strings`
	`sync`

	`github.com/mcuadros/go-defaults`
	`github.com/pelletier/go-toml`
	`github.com/storezhang/gox`
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
		c.path = flag.String("conf", "./conf/application.yaml", "指定配置文件路径")
		flag.StringVar(c.path, "c", *c.path, "指定配置文件路径")
		flag.Parse()
	})

	// 区分指针类型和非指针类型
	if reflect.ValueOf(config).Kind() == reflect.Ptr {
		err = c.loadConfig(config)
	} else {
		err = c.loadConfig(&config)
	}

	return
}

func (c *Config) loadConfig(config interface{}) (err error) {
	var finalPath string
	if finalPath, err = c.findConfigFilepath(*c.path); nil != err {
		return
	}

	// 可以处理后续动态加载
	if "" == c.format {
		c.format = strings.ToLower(filepath.Ext(finalPath))
	}
	if 0 == len(c.data) {
		if c.data, err = ioutil.ReadFile(finalPath); nil != err {
			return
		}
	}

	// 处理默认值
	if c.application.options.isDefault {
		defaults.SetDefaults(config)
	}

	switch c.format {
	case ".yml":
		fallthrough
	case ".yaml":
		err = yaml.Unmarshal(c.data, config)
	case ".json":
		err = json.Unmarshal(c.data, config)
	case ".toml":
		err = toml.Unmarshal(c.data, config)
	case ".xml":
		err = xml.Unmarshal(c.data, config)
	default:
		err = yaml.Unmarshal(c.data, config)
	}
	if nil != err {
		return
	}

	// 验证数据
	if c.application.options.isValidate {
		err = validatorx.Struct(config)
	}

	return
}

func (c *Config) findConfigFilepath(conf string) (path string, err error) {
	path = conf
	if "" != path && gox.IsFileExist(path) {
		return
	}

	var notExists bool
	if path, notExists = c.findConfigFilepathWithExt("./application"); !notExists {
		return
	}
	if path, notExists = c.findConfigFilepathWithExt("./conf/application"); !notExists {
		return
	}
	if path, notExists = c.findConfigFilepathWithExt("./app"); !notExists {
		return
	}
	if path, notExists = c.findConfigFilepathWithExt("./conf/app"); !notExists {
		return
	}
	err = errors.New("找不到配置文件")

	return
}

// 之所以命名为notExists，是为了少对notExists赋值
func (c *Config) findConfigFilepathWithExt(filename string) (path string, notExists bool) {
	path = fmt.Sprintf("%s.%s", filename, "yaml")
	if gox.IsFileExist(path) {
		return
	}

	path = fmt.Sprintf("%s.%s", filename, "yml")
	if gox.IsFileExist(path) {
		return
	}

	path = fmt.Sprintf("%s.%s", filename, "toml")
	if gox.IsFileExist(path) {
		return
	}

	path = fmt.Sprintf("%s.%s", filename, "json")
	if gox.IsFileExist(path) {
		return
	}

	path = fmt.Sprintf("%s.%s", filename, "xml")
	if gox.IsFileExist(path) {
		return
	}
	notExists = true

	return
}
