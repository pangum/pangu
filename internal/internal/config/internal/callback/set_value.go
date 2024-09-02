package callback

import (
	"reflect"

	"github.com/goexl/gox"
)

type SetValue func(names gox.Slice[string], field reflect.Value)
