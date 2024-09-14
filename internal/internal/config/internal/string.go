package internal

import (
	"reflect"
)

type String[F any] interface {
	String(from F, target reflect.Value) (err error)
}
