package arg

type base struct {
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
func DefaultArg(name string, usage string, aliases ...string) *base {
	return NewArg(name, usage, false, false, "", []string{}, aliases...)
}

// NewRequiredArg 创建必填参数
func NewRequiredArg(name string, usage string, aliases ...string) *base {
	return NewArg(name, usage, true, false, "", []string{}, aliases...)
}

// NewHiddenArg 创建隐藏参数
func NewHiddenArg(name string, usage string, aliases ...string) *base {
	return NewArg(name, usage, false, true, "", []string{}, aliases...)
}

// NewArg 创建参数
func NewArg(name string, usage string, required bool, hidden bool, defaultText string, envs []string, aliases ...string) *base {
	return &base{
		name:        name,
		aliases:     aliases,
		usage:       usage,
		envVars:     envs,
		required:    required,
		hidden:      hidden,
		defaultText: defaultText,
	}
}

func (a *base) Name() string {
	return a.name
}

func (a *base) Aliases() []string {
	return a.aliases
}

func (a *base) Usage() string {
	return a.usage
}

func (a *base) DefaultText() string {
	return a.defaultText
}
