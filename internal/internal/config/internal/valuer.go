package internal

import (
	"github.com/goexl/gox"
)

type Valuer[F any] interface {
	Int[F]
	Uint[F]
	Float[F]
	String[F]
	Bool[F]

	// Key 原始字段转换为新的字段
	Key(from gox.Slice[string]) (to gox.Slice[string])

	// Get 取出配置值
	Get(key string) (value F, ok bool)
}
