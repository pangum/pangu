package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewFloat64s
	_ app.Arg = (*float64sArg)(nil)
)

type float64sArg struct {
	*Argument

	destination *[]float64
}

// NewFloat64s 创建一个浮点数组参数
func NewFloat64s(name string, destination *[]float64, opts ...option) *float64sArg {
	return &float64sArg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (f *float64sArg) Destination() any {
	return f.destination
}

func (f *float64sArg) Flag() (flag app.Flag) {
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
