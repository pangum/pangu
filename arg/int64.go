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
	*Argument

	destination *int64
}

// NewInt64 创建一个整形参数
func NewInt64(name string, destination *int64, opts ...option) *int64Arg {
	return &int64Arg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (i *int64Arg) Destination() any {
	return i.destination
}

func (i *int64Arg) Flag() (flag app.Flag) {
	_if := &cli.Int64Flag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		DefaultText: i.DefaultText(),
		Required:    i.Required(),
		Hidden:      i.Hidden(),
	}
	if nil != i.Default() {
		_if.Value = i.Default().(int64)
	}
	if nil != i.Destination() {
		_if.Destination = i.Destination().(*int64)
	}
	flag = _if

	return
}
