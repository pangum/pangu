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

	destination *int
}

// NewInt 创建一个整形参数
func NewInt(name string, destination *int, opts ...option) *intArg {
	return &intArg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (i *intArg) Destination() any {
	return i.destination
}

func (i *intArg) Flag() (flag app.Flag) {
	_if := &cli.IntFlag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		DefaultText: i.DefaultText(),
		Required:    i.Required(),
		Hidden:      i.Hidden(),
	}
	if nil != i.Default() {
		_if.Value = i.Default().(int)
	}
	if nil != i.Destination() {
		_if.Destination = i.Destination().(*int)
	}
	flag = _if

	return
}
