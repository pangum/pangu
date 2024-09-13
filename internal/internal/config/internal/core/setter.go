package core

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/config/internal/callback"
	"github.com/pangum/pangu/internal/runtime"
)

type Setter struct {
	logger *log.Logger
}

func NewSetter(logger *log.Logger) *Setter {
	return &Setter{
		logger: logger,
	}
}

func (s *Setter) Process(target runtime.Pointer) {
	s.process(reflect.ValueOf(target).Elem(), gox.NewSlice[string](), s.setValue)
}

func (s *Setter) process(target reflect.Value, names gox.Slice[string], set callback.SetValue) {
	typ := target.Type()
	for index := 0; index < target.NumField(); index++ {
		target := target.Field(index)
		name := typ.Field(index).Name
		kind := target.Kind()

		if reflect.Struct == kind { // 如果字段为结构体才递归处理
			names = append(names, name)
			s.process(target, names, set)
		} else if reflect.Ptr == kind && target.Elem().Kind() == reflect.Struct { // 如果是结构体指针才进行递归处理
			if target.IsNil() { // 仅对空指针进行初始化
				target.Set(reflect.New(target.Type().Elem()))
			}
			names = append(names, name)
			s.process(target.Elem(), names, set)
		} else { // 不管是否被设置，都尝试从环境变量中设置变量值
			copies := names.Clone() // !复制一份字段名列表，防止干扰其它结构体字段
			set(append(copies, name), target)
		}
	}
}

func (s *Setter) setValue(names gox.Slice[string], target reflect.Value) {
	switch target.Kind() {
	case reflect.Bool:
		s.setEnvironmentValue(names, target, s.bool)
	case reflect.Int:
		s.setEnvironmentValue(names, target, s.int)
	case reflect.Int8:
		s.setEnvironmentValue(names, target, s.int8)
	case reflect.Int16:
		s.setEnvironmentValue(names, target, s.int16)
	case reflect.Int32:
		s.setEnvironmentValue(names, target, s.int32)
	case reflect.Int64:
		s.setEnvironmentValue(names, target, s.int64)
	case reflect.Uint:
		s.setEnvironmentValue(names, target, s.uint)
	case reflect.Uint8:
		s.setEnvironmentValue(names, target, s.uint8)
	case reflect.Uint16:
		s.setEnvironmentValue(names, target, s.uint16)
	case reflect.Uint32:
		s.setEnvironmentValue(names, target, s.uint32)
	case reflect.Uint64:
		s.setEnvironmentValue(names, target, s.uint64)
	case reflect.Uintptr:
		s.setEnvironmentValue(names, target, s.uintPtr)
	case reflect.Float32:
		s.setEnvironmentValue(names, target, s.float32)
	case reflect.Float64:
		s.setEnvironmentValue(names, target, s.float64)
	case reflect.String:
		s.setEnvironmentValue(names, target, s.string)
	default:
		s.setEnvironmentValue(names, target, s.string)
	}
}

func (s *Setter) setEnvironmentValue(names gox.Slice[string], target reflect.Value, convert callback.Convert) {
	// 将所有名称转换为大写，符合环境变量的定义
	keys := make([]string, names.Length()) // 复制一份，不影响原来的字段名
	for index, name := range names {
		keys[index] = strings.ToUpper(name)
	}

	key := strings.Join(keys, constant.Underline)
	fields := gox.Fields[any]{
		field.New("field", strings.Join(names, constant.Dot)),
		field.New("environment", key),
	}
	if environment, ok := os.LookupEnv(key); !ok {
		(*s.logger).Debug("未在环境变量中发现配置", fields...)
	} else if ce := convert(environment, target); nil != ce {
		(*s.logger).Debug("从环境变量中设置配置值出错", fields.Add(field.Error(ce))...)
	} else {
		(*s.logger).Debug("环境变量中设置配置值", fields...)
	}
}

func (s *Setter) bool(from string, target reflect.Value) (err error) {
	if value, pbe := strconv.ParseBool(from); nil == pbe {
		target.Set(reflect.ValueOf(value))
	} else {
		err = pbe
	}

	return
}

func (s *Setter) float32(from string, target reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 32); nil == pfe {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pfe
	}

	return
}

func (s *Setter) float64(from string, target reflect.Value) (err error) {
	if value, pfe := strconv.ParseFloat(from, 64); nil == pfe {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pfe
	}

	return
}

func (s *Setter) int(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) int8(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 8); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) int16(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 16); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) int32(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseInt(from, 0, 32); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) int64(from string, target reflect.Value) (err error) {
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
func (s *Setter) uint(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) uint8(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 8); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) uint16(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 16); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) uint32(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 32); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) uint64(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, 64); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) uintPtr(from string, target reflect.Value) (err error) {
	if value, pie := strconv.ParseUint(from, 0, strconv.IntSize); nil == pie {
		target.Set(reflect.ValueOf(value).Convert(target.Type()))
	} else {
		err = pie
	}

	return
}

func (s *Setter) string(from string, target reflect.Value) (err error) {
	target.Set(reflect.ValueOf(from).Convert(target.Type()))

	return
}
