package valuer

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/goexl/gox"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/config/internal"
)

var _ internal.Valuer[string] = (*Environment)(nil)

type Environment struct {
	// 无
}

func NewEnvironment() *Environment {
	return new(Environment)
}

func (Environment) Key(from gox.Slice[string]) gox.Slice[string] {
	for index, name := range from {
		from[index] = strings.ToUpper(name)
	}

	return gox.NewSlice(strings.Join(from, constant.Underline))
}

func (Environment) Get(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (Environment) Bool(from string, field reflect.Value) (err error) {
	if value, pbe := strconv.ParseBool(from); nil == pbe {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pbe
	}

	return
}

func (Environment) Float32(from string, target reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 32); nil == pfe {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pfe
	}

	return
}

func (Environment) Float64(from string, target reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 64); nil == pfe {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pfe
	}

	return
}

func (Environment) Int(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) Int8(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 8); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) Int16(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 16); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) Int32(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 32); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) Int64(from string, target reflect.Value) (err error) {
	var value any
	switch target.Interface().(type) {
	case time.Duration:
		value, err = time.ParseDuration(from)
	case gox.Bytes:
		value, err = gox.ParseBytes(from)
	default:
		value, err = strconv.ParseInt(from, 0, 64)
	}
	if nil != err {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	}

	return
}

func (Environment) Uint(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) Uint8(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 8); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) Uint16(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 16); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) Uint32(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 32); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) Uint64(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 64); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) Uintptr(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (Environment) String(from string, target reflect.Value) (err error) {
	target.Set(reflect.ValueOf(from).Convert(target.Type()))

	return
}
