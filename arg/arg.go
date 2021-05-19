package arg

type arg struct {
	// 名称
	name string
	// 别名
	aliases []string
	// 使用方法
	usage string
	// 环境变量
	envVars []string
	// 是否是必须
	required bool
	// 是否隐藏
	hidden bool
	// 默认值
	defaultText string
}

// DefaultArg 创建默认参数
func DefaultArg(name string, usage string, aliases ...string) arg {
	return NewArg(name, usage, false, false, "", []string{}, aliases...)
}

// NewArg 创建参数
func NewArg(name string, usage string, required bool, hidden bool, defaultText string, envs []string, aliases ...string) arg {
	return arg{
		name:        name,
		aliases:     aliases,
		usage:       usage,
		envVars:     envs,
		required:    required,
		hidden:      hidden,
		defaultText: defaultText,
	}
}

func (a *arg) Name() string {
	return a.name
}

func (a *arg) Aliases() []string {
	return a.aliases
}

func (a *arg) Usage() string {
	return a.usage
}

func (a *arg) DefaultText() string {
	return a.defaultText
}
