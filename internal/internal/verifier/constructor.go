package verifier

import (
	"reflect"
	"runtime"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/internal/message"
	"github.com/pangum/pangu/internal/param"
)

type Constructor struct {
	params *param.Application
}

func NewConstructor(params *param.Application) *Constructor {
	return &Constructor{
		params: params,
	}
}

func (c *Constructor) Verify(constructor any) (err error) {
	if c.params.Verify {
		return
	}

	constructorType := reflect.TypeOf(constructor)
	// 构造方法必须是方法不能是其它类型
	if reflect.Func != constructorType.Kind() {
		err = exc.NewField(message.ConstructorMustFunc, field.New("constructor", constructorType.String()))
	}
	if nil != err {
		return
	}

	// 构造方法必须有返回值
	if 0 == constructorType.NumOut() {
		constructorName := runtime.FuncForPC(reflect.ValueOf(constructor).Pointer()).Name()
		err = exc.NewField(message.ConstructorMustReturn, field.New("constructor", constructorName))
	}

	return
}
