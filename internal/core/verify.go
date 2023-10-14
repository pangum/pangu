package core

import (
	"reflect"
	"runtime"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/exception"
	"github.com/pangum/pangu/internal/param"
)

type Verify struct {
	param *param.Application
}

func (v *Verify) Bootstrap(bootstrap any) (err error) {
	if v.param.Verify {
		return
	}

	constructorType := reflect.TypeOf(bootstrap)
	// 构造方法必须是方法不能是其它类型
	if reflect.Func != constructorType.Kind() {
		err = exc.NewField(exception.ConstructorMustFunc, field.New("bootstrap", constructorType.String()))
	}
	if nil != err {
		return
	}

	// 构造方法必须有依赖项
	if 0 == constructorType.NumIn() {
		constructorName := runtime.FuncForPC(reflect.ValueOf(bootstrap).Pointer()).Name()
		err = exc.NewField(exception.BootstrapMustHasDependencies, field.New("bootstrap", constructorName))
	}
	if nil != err {
		return
	}

	// 只能返回一个类型为Bootstrap返回值
	returnsCount := constructorType.NumOut()
	if 1 != returnsCount || reflect.TypeOf((*Bootstrap)(nil)).Elem() != constructorType.Out(constant.IndexFirst) {
		err = exc.NewMessage(exception.BootstrapMustReturnBootstrap)
	}

	return
}

func (v *Verify) Constructor(constructor any) (err error) {
	if v.param.Verify {
		return
	}

	constructorType := reflect.TypeOf(constructor)
	// 构造方法必须是方法不能是其它类型
	if reflect.Func != constructorType.Kind() {
		err = exc.NewField(exception.ConstructorMustFunc, field.New("constructor", constructorType.String()))
	}
	if nil != err {
		return
	}

	// 构造方法必须有返回值
	if 0 == constructorType.NumOut() {
		constructorName := runtime.FuncForPC(reflect.ValueOf(constructor).Pointer()).Name()
		err = exc.NewField(exception.ConstructorMustReturn, field.New("constructor", constructorName))
	}

	return
}
