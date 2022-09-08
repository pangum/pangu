package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewBool
	_ app.Arg = (*boolArg)(nil)
)

type boolArg struct {
	*Argument
}

// NewBool 创建布尔参数
func NewBool(name string, opts ...option) *boolArg {
	return &boolArg{
		Argument: New(name, opts...),
	}
}

func (b *boolArg) Flag() app.Flag {
	return &cli.BoolFlag{
		Name:        b.Name(),
		Aliases:     b.Aliases(),
		Usage:       b.Usage(),
		Destination: b.Destination().(*bool),
		Value:       b.Default().(bool),
		DefaultText: b.DefaultText(),
		Required:    b.Required(),
		Hidden:      b.Hidden(),
	}
}
