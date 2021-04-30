package app

type Arg interface {
	flag

	// GetValue 参数值
	GetValue() interface{}
	// GetDefaultText 默认值
	GetDefaultText() string
}
