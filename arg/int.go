package arg

import (
	"github.com/urfave/cli/v2"

	"github.com/pangum/pangu/app"
)

var (
	_         = NewInt
	_ app.Arg = (*intArg)(nil)
)

type intArg struct {
	*Argument
}

// NewInt 创建一个整形参数
func NewInt(name string, opts ...option) *intArg {
	return &intArg{
		Argument: New(name, opts...),
	}
}

func (i *intArg) Flag() app.Flag {
	return &cli.IntFlag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Destination: i.Destination().(*int),
		Value:       i.Default().(int),
		DefaultText: i.DefaultText(),
		Required:    i.Required(),
		Hidden:      i.Hidden(),
	}
}
