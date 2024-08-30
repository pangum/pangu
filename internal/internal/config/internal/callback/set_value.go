package callback

import (
	"reflect"
)

type SetValue func(names []string, field reflect.Value)
