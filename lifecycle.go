package core

import (
	"context"
)

// Lifecycle 生命周期方法
type Lifecycle struct {
	// 用于编写代码时从本结构体继承（匿名内部类）从而少写部分方法
}

func (l *Lifecycle) Before(_ context.Context) (err error) {
	return
}

func (l *Lifecycle) After(_ context.Context) (err error) {
	return
}
