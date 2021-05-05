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
	`github.com/storezhang/pangu/app`
	`github.com/storezhang/pangu/command`
	`github.com/urfave/cli/v2`
	`gopkg.in/yaml.v3`
)

type Bootstrap struct {
	application Application
	// 内置命令
	serve   *command.Serve
	version *command.Version

	// 存储配置
	config interface{}

	// 应用程序本身的配置
	// 徽标
	banner string

	cli *cli.App
}

// NewBootstrap 创建启动器
func NewBootstrap(
	application Application,
	serve *command.Serve, version *command.Version,
	cli *cli.App,
) Bootstrap {
	return Bootstrap{
		application: application,
		serve:       serve,
		version:     version,

		cli: cli,
	}
}

// Start 开始执行
func (b *Bootstrap) Start() error {
	// 添加内置命令
	b.addInternalFlags()
	b.addInternalCommands()

	// 指定初始行为是加载配置文件
	b.cli.Action = b.loadConfig
	// 应用程序

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

// SetBanner 设置徽标
func (b *Bootstrap) SetBanner(banner string) {
	b.banner = banner
}

func (b *Bootstrap) addInternalCommands() {
	b.cli.Commands = append(b.cli.Commands, &cli.Command{
		Name:    b.serve.GetName(),
		Aliases: b.serve.GetAliases(),
		Usage:   b.serve.GetUsage(),
		Action: func(ctx *cli.Context) error {
			return b.serve.Run(app.NewContext(ctx))
		},
	})
	b.cli.Commands = append(b.cli.Commands, &cli.Command{
		Name:    b.version.GetName(),
		Aliases: b.version.GetAliases(),
		Usage:   b.version.GetUsage(),
		Action: func(ctx *cli.Context) error {
			return b.version.Run(app.NewContext(ctx))
		},
	})
}

func (b *Bootstrap) addInternalFlags() {
	b.cli.Flags = append(b.cli.Flags, &cli.StringFlag{
		Name:        "conf",
		Aliases:     []string{"c", "config", "configuration"},
		Usage:       "",
		Value:       "./conf/application.yaml",
		DefaultText: "./conf/application.yaml",
	})
}

func (b *Bootstrap) parseArgs() (flags []cli.Flag) {
	// 添加内置参数

	return
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
