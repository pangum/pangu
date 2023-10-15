package config

import (
	"github.com/pangum/pangu/internal/runtime"
)

type Loader interface {
	// Load 加载配置
	Load(path string, config runtime.Pointer) (err error)
}
