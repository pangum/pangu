package param

import (
	"time"
)

type Timeout struct {
	// 启动
	Boot time.Duration
	// 退出
	Quit time.Duration
	Set  bool
}

func NewTimeout() *Timeout {
	return &Timeout{
		Boot: 15 * time.Second,
		Quit: 15 * time.Second,
	}
}
