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
	*base
}

// NewFloat64s 创建一个浮点数组参数
func NewFloat64s(name string, opts ...option) *float64sArg {
	return &float64sArg{
		base: _new(name, opts...),
	}
}
func (f *float64sArg) Flag() app.Flag {
	return &cli.Float64SliceFlag{
		Name:        f.Name(),
		Aliases:     f.Aliases(),
		Usage:       f.Usage(),
		Destination: cli.NewFloat64Slice(*f.Destination().(*[]float64)...),
		Value:       cli.NewFloat64Slice(f.Default().([]float64)...),
		DefaultText: f.DefaultText(),
		Required:    f.Required(),
		Hidden:      f.Hidden(),
	}
}
