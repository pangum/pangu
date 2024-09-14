package internal

import (
	"reflect"
)

type Int[F any] interface {
	Int(from F, target reflect.Value) (err error)

	Int8(from F, target reflect.Value) (err error)

	Int16(from F, target reflect.Value) (err error)

	Int32(from F, target reflect.Value) (err error)

	Int64(from F, target reflect.Value) (err error)
}
