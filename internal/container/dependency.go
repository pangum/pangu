package container

import (
	"reflect"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/pangum/pangu/internal/internal/message"
	"github.com/pangum/pangu/internal/internal/param"
	"github.com/pangum/pangu/internal/runtime"
	"go.uber.org/dig"
)

type Dependency struct {
	container *dig.Container
	params    *param.Dependency
}

func NewDependency(container *dig.Container, params *param.Dependency) *Dependency {
	return &Dependency{
		container: container,
		params:    params,
	}
}

func (d *Dependency) Inject() (err error) {
	if de := d.puts(); nil != de {
		err = de
	} else if ge := d.gets(); nil != ge {
		err = ge
	} else {
		d.params.Clear() // 确保连续调用时，不会有脏数据
	}

	return
}

func (d *Dependency) Apply() {
	if err := d.Inject(); nil != err {
		panic(err)
	}
}

func (d *Dependency) gets() (err error) {
	for _, get := range d.params.Gets {
		err = d.invoke(get)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependency) invoke(get *param.Get) error {
	return d.container.Invoke(get.Getter)
}

func (d *Dependency) puts() (err error) {
	for _, put := range d.params.Puts {
		err = d.put(put)
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependency) put(put *param.Put) error {
	options := make([]dig.ProvideOption, 0)
	if "" != put.Name {
		options = append(options, dig.Name(put.Name))
	}
	if "" != put.Group {
		options = append(options, dig.Group(put.Group))
	}

	return d.provide(put.Constructor, options...)
}

func (d *Dependency) provide(constructor runtime.Constructor, options ...dig.ProvideOption) (err error) {
	if ve := d.verify(constructor); nil != ve {
		err = ve
	} else if pe := d.container.Provide(constructor, options...); nil != pe {
		err = pe
	}

	return
}

func (d *Dependency) verify(constructor any) (err error) {
	if d.params.Verify {
		return
	}

	typ := reflect.TypeOf(constructor)
	if reflect.Func != typ.Kind() { // 构造方法必须是方法，不能是其它类型
		err = exception.New().Message(message.ConstructorMustFunc).Field(field.New("constructor", typ.String())).Build()
	} else if 0 == typ.NumOut() { // 构造方法必须有返回值
		name := runtime.FuncForPC(reflect.ValueOf(constructor).Pointer()).Name()
		err = exception.New().Message(message.ConstructorMustReturn).Field(field.New("constructor", name)).Build()
	}

	return
}
