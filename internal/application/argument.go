package application

// Argument 参数
type Argument interface {
	Parameter

	// Target 绑定值
	Target() any
}
