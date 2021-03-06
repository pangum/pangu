package pangu

const (
	defaultName = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Name=$NAME"来注入值`
	defaultsTag = `default`
	firstIndex  = 0

	copyright   = `https://pangu.pangum.tech`
	authorName  = `storezhang`
	authorEmail = `storezhang@gmail.com`

	configLongName  = `conf`
	configShortName = `c`
	configUsage     = `指定配置文件路径`
	configDefault   = `./conf/application.yaml`
	configDir       = `config`
	confDir         = `conf`
	applicationName = `application`

	envSeparator = `=`
	envCount     = 2

	yamlExt = `.yaml`
	ymlExt  = `.yml`
	jsonExt = `.json`
	xmlExt  = `.xml`
	tomlExt = `.toml`
)
