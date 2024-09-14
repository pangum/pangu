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

var _ internal.Valuer[string] = (*String)(nil)

type String struct {
	// 无
}

func NewString() *String {
	return new(String)
}

func (String) Key(from gox.Slice[string]) gox.Slice[string] {
	for index, name := range from {
		from[index] = strings.ToUpper(name)
	}

	return gox.NewSlice(strings.Join(from, constant.Underline))
}

func (String) Get(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (String) Bool(from string, field reflect.Value) (err error) {
	if value, pbe := strconv.ParseBool(from); nil == pbe {
		field.Set(reflect.ValueOf(value))
	} else {
		err = pbe
	}

	return
}

func (String) Float32(from string, target reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 32); nil == pfe {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pfe
	}

	return
}

func (String) Float64(from string, target reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 64); nil == pfe {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pfe
	}

	return
}

func (String) Int(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) Int8(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 8); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) Int16(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 16); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) Int32(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 32); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) Int64(from string, target reflect.Value) (err error) {
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

func (String) Uint(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) Uint8(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 8); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) Uint16(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 16); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) Uint32(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 32); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) Uint64(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 64); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) Uintptr(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (String) String(from string, target reflect.Value) (err error) {
	target.Set(reflect.ValueOf(from).Convert(target.Type()))

	return
}
