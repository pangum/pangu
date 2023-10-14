package app

// Argument 参数
type Argument interface {
	Parameter

	// Default 参数默认值
	Default() any

	// Target 绑定值
	Target() any

	// Text 默认显示文本
	Text() string

	// Required 是否是必填项
	Required() bool

	// Flag 解析出真正使用的参数
	Flag() Flag
}
