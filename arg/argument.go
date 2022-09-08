package arg

import (
	"github.com/goexl/gox"
)

// Argument 参数
type Argument struct {
	// 名称
	name string
	// 默认值
	_default any
	// 绑定
	destination any
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
	// 默认显示字符串
	dt string

	_ gox.CannotCopy
}

func New(name string, opts ...option) *Argument {
	_options := defaultOption()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return &Argument{
		name:        name,
		_default:    _options._default,
		destination: _options.destination,
		aliases:     _options.aliases,
		usage:       _options.usage,
		envs:        _options.envs,
		required:    _options.required,
		hidden:      _options.hidden,
		dt:          _options.dt,
	}
}

func (b *Argument) Name() string {
	return b.name
}

func (b *Argument) Default() any {
	return b._default
}

func (b *Argument) Destination() any {
	return b.destination
}

func (b *Argument) Aliases() []string {
	return b.aliases
}

func (b *Argument) Usage() string {
	return b.usage
}

func (b *Argument) DefaultText() string {
	return b.dt
}

func (b *Argument) Required() bool {
	return b.required
}

func (b *Argument) Hidden() bool {
	return b.hidden
}
