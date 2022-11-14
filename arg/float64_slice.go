package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewFloat64s
	_ app.Arg = (*float64SliceArg)(nil)
)

type float64SliceArg struct {
	*Argument

	destination *[]float64
}

// NewFloat64s 创建一个浮点数组参数
func NewFloat64s(name string, destination *[]float64, opts ...option) *float64SliceArg {
	return &float64SliceArg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (f *float64SliceArg) Destination() any {
	return f.destination
}

func (f *float64SliceArg) Flag() (flag app.Flag) {
	ff := &cli.Float64SliceFlag{
		Name:        f.Name(),
		Aliases:     f.Aliases(),
		Usage:       f.Usage(),
		DefaultText: f.DefaultText(),
		Required:    f.Required(),
		Hidden:      f.Hidden(),
	}
	if nil != f.Default() {
		ff.Value = cli.NewFloat64Slice(f.Default().([]float64)...)
	}
	if nil != f.Destination() {
		ff.Destination = cli.NewFloat64Slice(*f.Destination().(*[]float64)...)
	}
	flag = ff

	return
}
