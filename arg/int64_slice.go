package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewInt64s
	_ app.Arg = (*int64sArg)(nil)
)

type int64sArg struct {
	*Argument

	destination []int64
}

// NewInt64s 创建一个整形数组参数
func NewInt64s(name string, destination []int64, opts ...option) *int64sArg {
	return &int64sArg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (i *int64sArg) Destination() any {
	return i.destination
}

func (i *int64sArg) Flag() (flag app.Flag) {
	isf := &cli.Int64SliceFlag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		DefaultText: i.DefaultText(),
		Required:    i.Required(),
		Hidden:      i.Hidden(),
	}
	if nil != i.Default() {
		isf.Value = cli.NewInt64Slice(i.Default().([]int64)...)
	}
	if nil != i.Destination() {
		isf.Destination = cli.NewInt64Slice(*i.Destination().(*[]int64)...)
	}
	flag = isf

	return
}
