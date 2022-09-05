package pangu

type configLoader interface {
	// Load 加载配置
	Load(path string, config any) (err error)
}
