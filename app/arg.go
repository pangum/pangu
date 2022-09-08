package app

type Arg interface {
	parameter

	// Value 参数值
	Value() interface{}

	// Default 默认值
	Default() string

	// Flag 解析出真正使用的参数
	Flag() Flag
}
