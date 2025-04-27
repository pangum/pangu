package config

import (
	"time"
)

type Timeout struct {
	// 启动
	Startup time.Duration
	// 退出
	Quit time.Duration
}

func NewTimeout() *Timeout {
	return &Timeout{
		Startup: 15 * time.Second,
		Quit:    15 * time.Second,
	}
}
