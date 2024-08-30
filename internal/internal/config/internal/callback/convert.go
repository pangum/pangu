package callback

import (
	"reflect"
)

type Convert func(from string, field reflect.Value) (err error)
