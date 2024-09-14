package internal

import (
	"reflect"
)

type Bool[F any] interface {
	Bool(from F, target reflect.Value) (err error)
}
