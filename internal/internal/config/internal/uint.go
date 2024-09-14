package internal

import (
	"reflect"
)

type Uint[F any] interface {
	Uint(from F, target reflect.Value) (err error)

	Uint8(from F, target reflect.Value) (err error)

	Uint16(from F, target reflect.Value) (err error)

	Uint32(from F, target reflect.Value) (err error)

	Uint64(from F, target reflect.Value) (err error)

	Uintptr(from F, target reflect.Value) (err error)
}
