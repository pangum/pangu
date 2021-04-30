package app

type flag interface {
	// GetAliases 别名，可以是短名称也可以是长名称，比如一个叫version的命令，别名可以是[v,V,Version]
	GetAliases() []string

	// GetName 名称
	GetName() string

	// GetUsage 描述使用方法
	GetUsage() string
}
