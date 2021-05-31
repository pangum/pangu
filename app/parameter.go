package app

type parameter interface {
	// Aliases 别名，可以是短名称也可以是长名称，比如一个叫version的命令，别名可以是[v,version]
	Aliases() []string

	// Name 名称
	Name() string

	// Usage 描述使用方法
	Usage() string
}
