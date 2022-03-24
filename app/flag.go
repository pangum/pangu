package app

import (
	`flag`
	`fmt`
)

// Flag 描述一个可以被解析的参数或者选项
type Flag interface {
	fmt.Stringer

	// Apply 解析参数
	Apply(set *flag.FlagSet) error

	// Names 别名列表
	Names() []string

	// IsSet 是否可设置
	IsSet() bool
}
