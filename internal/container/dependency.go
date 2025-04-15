package container

import (
	"reflect"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/pangum/core/internal/internal/constant"
	"github.com/pangum/core/internal/internal/message"
	"github.com/pangum/core/internal/internal/param"
	"github.com/pangum/core/internal/runtime"
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
		if ve := d.verify(put.Constructor); nil != ve {
			err = ve
		} else if pe := d.put(put); nil != pe {
			err = pe
		}

		if nil != err {
			break
		}
	}

	return
}

func (d *Dependency) put(put *param.Put) (err error) {
	for _, name := range put.Names {
		for _, group := range put.Groups {
			err = d.provide(put.Constructor, name, group)
		}
		if nil != err {
			break
		}
	}

	return
}

func (d *Dependency) provide(constructor any, name string, group string) error {
	options := make([]dig.ProvideOption, 0)
	if constant.DependencyNone != name {
		options = append(options, dig.Name(name))
	}
	if constant.DependencyNone != group {
		options = append(options, dig.Group(group))
	}

	return d.container.Provide(constructor, options...)
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
