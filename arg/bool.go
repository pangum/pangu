package arg

import (
	`github.com/storezhang/pangu/app`
)

var _ app.Arg = (*Bool)(nil)

// Bool 布尔值参数
type Bool struct {
	Arg

	// 值
	Value bool
}

func (b *Bool) GetValue() interface{} {
	return b.Value
}
