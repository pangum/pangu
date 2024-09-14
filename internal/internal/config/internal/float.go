package internal

import (
	"reflect"
)

type Float[F any] interface {
	Float32(from F, target reflect.Value) (err error)

	Float64(from F, target reflect.Value) (err error)
}
