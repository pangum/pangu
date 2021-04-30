package pangu

import (
	`encoding/json`
	`errors`
	`io/ioutil`
	`os`
	`path/filepath`
	`strings`

	`github.com/pelletier/go-toml`
	`github.com/storezhang/gox`
	`github.com/urfave/cli/v2`
	`gopkg.in/yaml.v3`
)

type Bootstrap struct {
	// 存储配置
	config interface{}

	cli *cli.App
}

// NewBootstrap 创建启动器
func NewBootstrap(cli *cli.App) Bootstrap {
	return Bootstrap{
		cli: cli,
	}
}

// Start 开始执行
func (b *Bootstrap) Start() error {
	b.cli.Action = b.loadConfig

	return b.cli.Run(os.Args)
}

// SetConfig 添加一个服务器到应用程序中
func (b *Bootstrap) SetConfig(config interface{}) {
	b.config = config
}

// GetConfig 取得解析后的配置
func (b *Bootstrap) GetConfig() interface{} {
	return b.config
}

func (b *Bootstrap) loadConfig(ctx *cli.Context) (err error) {
	var conf string
	if conf, err = b.findConfigFilepath(ctx.String("conf")); nil != err {
		return
	}

	var data []byte
	if data, err = ioutil.ReadFile(conf); nil != err {
		return
	}

	switch strings.ToLower(filepath.Ext(conf)) {
	case "yml":
	case "yaml":
		err = yaml.Unmarshal(data, &b.config)
	case "json":
		err = json.Unmarshal(data, &b.config)
	case "toml":
		err = toml.Unmarshal(data, &b.config)
	default:
		err = yaml.Unmarshal(data, &b.config)
	}

	return
}

func (b *Bootstrap) findConfigFilepath(conf string) (filepath string, err error) {
	filepath = conf
	if "" == filepath {
		filepath = "./application.yml"
	}
	if gox.IsFileExist(filepath) {
		return
	}

	filepath = "./application.yaml"
	if gox.IsFileExist(filepath) {
		return
	}

	filepath = "./application.toml"
	if gox.IsFileExist(filepath) {
		return
	}

	filepath = "./conf/application.yml"
	if gox.IsFileExist(filepath) {
		return
	}

	filepath = "./conf/application.yaml"
	if gox.IsFileExist(filepath) {
		return
	}
	err = errors.New("找不到配置文件")

	return
}
