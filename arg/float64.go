package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewFloat64
	_ app.Arg = (*float64Arg)(nil)
)

type float64Arg struct {
	*Argument

	destination *float64
}

// NewFloat64 创建一个浮点型参数
func NewFloat64(name string, destination *float64, opts ...option) *float64Arg {
	return &float64Arg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (f *float64Arg) Destination() any {
	return f.destination
}

func (f *float64Arg) Flag() (flag app.Flag) {
	ff := &cli.Float64Flag{
		Name:        f.Name(),
		Aliases:     f.Aliases(),
		Usage:       f.Usage(),
		DefaultText: f.DefaultText(),
		Required:    f.Required(),
		Hidden:      f.Hidden(),
	}
	if nil != f.Default() {
		ff.Value = f.Default().(float64)
	}
	if nil != f.Destination() {
		ff.Destination = f.Destination().(*float64)
	}
	flag = ff

	return
}
