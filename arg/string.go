package arg

import (
	`github.com/storezhang/pangu/app`
)

var _ app.Arg = (*String)(nil)

// String 描述一个字符串参数
type String struct {
	Arg

	// 值
	Value string
}

func (s *String) GetValue() interface{} {
	return s.Value
}
