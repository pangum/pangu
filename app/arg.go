package app

type Arg interface {
	parameter

	// Value 参数值
	Value() interface{}

	// DefaultText 默认值
	DefaultText() string

	// ParseFlag 解析出真正使用的参数
	ParseFlag() Flag
}
