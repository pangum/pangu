package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewInt64
	_ app.Arg = (*int64Arg)(nil)
)

type int64Arg struct {
	*base
}

// NewInt64 创建一个整形参数
func NewInt64(name string, opts ...option) *int64Arg {
	return &int64Arg{
		base: _new(name, opts...),
	}
}

func (i *int64Arg) Default() interface{} {
	return i.Default
}

func (i *int64Arg) Flag() app.Flag {
	return &cli.Int64Flag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Destination: i.Destination().(*int64),
		Value:       i.Default().(int64),
		DefaultText: i.DefaultText(),
		Required:    i.Required(),
		Hidden:      i.Hidden(),
	}
}
