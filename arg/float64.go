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
}

// NewFloat64 创建一个浮点型参数
func NewFloat64(name string, opts ...option) *float64Arg {
	return &float64Arg{
		Argument: New(name, opts...),
	}
}

func (f *float64Arg) Flag() app.Flag {
	return &cli.Float64Flag{
		Name:        f.Name(),
		Aliases:     f.Aliases(),
		Usage:       f.Usage(),
		Destination: f.Destination().(*float64),
		Value:       f.Default().(float64),
		DefaultText: f.DefaultText(),
		Required:    f.Required(),
		Hidden:      f.Hidden(),
	}
}
