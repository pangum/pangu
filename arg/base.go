package arg

var (
	_ = DefaultArg
	_ = NewRequiredArg
	_ = NewHiddenArg
)

type base struct {
	// 名称
	name string
	// 别名
	aliases []string
	// 使用方法
	usage string
	// 环境变量列表
	envs []string
	// 是否是必须
	required bool
	// 是否隐藏
	hidden bool
	// 默认值
	defaultText string
}

// DefaultArg 创建默认参数
func DefaultArg(name string, usage string, aliases ...string) *base {
	return NewArg(name, usage, false, false, ``, []string{}, aliases...)
}

// NewRequiredArg 创建必填参数
func NewRequiredArg(name string, usage string, aliases ...string) *base {
	return NewArg(name, usage, true, false, ``, []string{}, aliases...)
}

// NewHiddenArg 创建隐藏参数
func NewHiddenArg(name string, usage string, aliases ...string) *base {
	return NewArg(name, usage, false, true, ``, []string{}, aliases...)
}

// NewArg 创建参数
func NewArg(
	name string, usage string,
	required bool, hidden bool,
	text string,
	envs []string, aliases ...string,
) *base {
	return &base{
		name:        name,
		aliases:     aliases,
		usage:       usage,
		envs:        envs,
		required:    required,
		hidden:      hidden,
		defaultText: text,
	}
}

func (b *base) Name() string {
	return b.name
}

func (b *base) Aliases() []string {
	return b.aliases
}

func (b *base) Usage() string {
	return b.usage
}

func (b *base) Default() string {
	return b.defaultText
}
