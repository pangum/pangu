package verifier

import (
	"reflect"
	"runtime"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/message"
	"github.com/pangum/pangu/internal/param"
	runtime2 "github.com/pangum/pangu/internal/runtime"
)

type Bootstrap struct {
	params *param.Application
}

func NewBoostrap(params *param.Application) *Bootstrap {
	return &Bootstrap{
		params: params,
	}
}

func (b *Bootstrap) Verify(bootstrap any) (err error) {
	if b.params.Verify {
		return
	}

	constructorType := reflect.TypeOf(bootstrap)
	// 构造方法必须是方法不能是其它类型
	if reflect.Func != constructorType.Kind() {
		err = exc.NewField(message.ConstructorMustFunc, field.New("bootstrap", constructorType.String()))
	}
	if nil != err {
		return
	}

	// 构造方法必须有依赖项
	if 0 == constructorType.NumIn() {
		constructorName := runtime.FuncForPC(reflect.ValueOf(bootstrap).Pointer()).Name()
		err = exc.NewField(message.BootstrapMustHasDependencies, field.New("bootstrap", constructorName))
	}
	if nil != err {
		return
	}

	// 只能返回一个类型为Bootstrap返回值
	returnsCount := constructorType.NumOut()
	if 1 != returnsCount || reflect.TypeOf((*runtime2.Bootstrap)(nil)).Elem() != constructorType.Out(constant.IndexFirst) {
		err = exc.NewMessage(message.BootstrapMustReturnBootstrap)
	}

	return
}
