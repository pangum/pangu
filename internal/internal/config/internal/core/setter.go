package core

import (
	"reflect"
	"strings"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/config/internal"
	"github.com/pangum/pangu/internal/internal/config/internal/callback"
	"github.com/pangum/pangu/internal/runtime"
)

type Setter[F any] struct {
	valuer internal.Valuer[F]
	logger *log.Logger
}

func NewSetter[F any](valuer internal.Valuer[F], logger *log.Logger) *Setter[F] {
	return &Setter[F]{
		valuer: valuer,
		logger: logger,
	}
}

func (s *Setter[F]) Process(config runtime.Pointer) {
	s.process(reflect.ValueOf(config).Elem(), gox.NewSlice[string](), s.set)
}

func (s *Setter[F]) process(config reflect.Value, names gox.Slice[string], set callback.SetValue) {
	typ := config.Type()
	for index := 0; index < config.NumField(); index++ {
		target := config.Field(index)
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

func (s *Setter[F]) set(names gox.Slice[string], target reflect.Value) {
	keys := s.valuer.Key(names.Clone())
	for _, key := range keys {
		fields := gox.Fields[any]{
			field.New("field", strings.Join(names, constant.Dot)),
			field.New("value", key),
		}
		if value, ok := s.valuer.Get(key); !ok {
			(*s.logger).Debug("未发现配置", fields...)
		} else if ce := s.setValue(value, target); nil != ce {
			(*s.logger).Debug("设置配置值出错", fields.Add(field.Error(ce))...)
		} else {
			(*s.logger).Debug("设置配置值成功", fields...)
		}
	}
}

func (s *Setter[F]) setValue(value F, target reflect.Value) (err error) {
	switch target.Kind() {
	case reflect.Bool:
		err = s.valuer.Bool(value, target)
	case reflect.Int:
		err = s.valuer.Int(value, target)
	case reflect.Int8:
		err = s.valuer.Int8(value, target)
	case reflect.Int16:
		err = s.valuer.Int16(value, target)
	case reflect.Int32:
		err = s.valuer.Int32(value, target)
	case reflect.Int64:
		err = s.valuer.Int64(value, target)
	case reflect.Uint:
		err = s.valuer.Uint(value, target)
	case reflect.Uint8:
		err = s.valuer.Uint8(value, target)
	case reflect.Uint16:
		err = s.valuer.Uint16(value, target)
	case reflect.Uint32:
		err = s.valuer.Uint32(value, target)
	case reflect.Uint64:
		err = s.valuer.Uint64(value, target)
	case reflect.Uintptr:
		err = s.valuer.Uintptr(value, target)
	case reflect.Float32:
		err = s.valuer.Float32(value, target)
	case reflect.Float64:
		err = s.valuer.Float64(value, target)
	case reflect.String:
		err = s.valuer.String(value, target)
	default:
		err = s.valuer.String(value, target)
	}

	return
}
