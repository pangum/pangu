package valuer

import (
	"reflect"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/config/internal"
)

var _ internal.Valuer[any] = (*Map)(nil)

type Map struct {
	original map[string]any
	str      *String
}

func NewMap(original map[string]any) *Map {
	return &Map{
		original: original,
		str:      NewString(),
	}
}

func (m Map) Key(from gox.Slice[string]) gox.Slice[string] {
	for index, name := range from {
		from[index] = strings.ToUpper(name)
	}

	return gox.NewSlice(strings.Join(from, constant.Underline))
}

func (m Map) Get(key string) (value any, ok bool) {
	if value, ok = m.original[key]; ok {
		// !没有任何操作，纯粹是语言上的缺陷
	}

	return
}

func (m Map) Bool(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case bool:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Bool(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Float32(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case float32:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Float32(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Float64(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case float64:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Float64(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Int(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case int:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Int(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Int8(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case int8:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Int8(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Int16(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case int16:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Int16(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Int32(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case int32:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Int32(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Int64(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case int64:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Int64(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Uint(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case uint:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Uint(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Uint8(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case uint8:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Uint8(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Uint16(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case uint16:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Uint16(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Uint32(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case uint32:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Uint32(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Uint64(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case uint64:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Uint64(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) Uintptr(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case *uint:
		target.Set(reflect.ValueOf(typed))
	case string:
		err = m.str.Uintptr(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}

func (m Map) String(from any, target reflect.Value) (err error) {
	switch typed := from.(type) {
	case string:
		err = m.str.String(typed, target)
	default:
		err = exception.New().Message("不被支持的配置项").Field(field.New("from", from)).Build()
	}

	return
}
