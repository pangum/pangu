package app

type Arg interface {
	parameter

	// Default 参数默认值
	Default() any

	// Destination 绑定值
	Destination() any

	// DefaultText 默认显示文本
	DefaultText() string

	// Required 是否是必填项
	Required() bool

	// Hidden 是否是隐藏项
	Hidden() bool

	// Flag 解析出真正使用的参数
	Flag() Flag
}
