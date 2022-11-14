package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewIntSlice
	_ app.Arg = (*intSliceArg)(nil)
)

type intSliceArg struct {
	*Argument

	destination *[]int
}

// NewIntSlice 创建一个整形数组参数
func NewIntSlice(base *Argument, destination *[]int, values ...int) *intSliceArg {
	return &intSliceArg{
		Argument:    base,
		destination: destination,
	}
}

func (i *intSliceArg) Destination() any {
	return i.destination
}

func (i *intSliceArg) Flag() (flag app.Flag) {
	isf := &cli.IntSliceFlag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		DefaultText: i.DefaultText(),
		Required:    i.Required(),
		Hidden:      i.Hidden(),
	}
	if nil != i.Default() {
		isf.Value = cli.NewIntSlice(i.Default().([]int)...)
	}
	if nil != i.Destination() {
		isf.Destination = cli.NewIntSlice(*i.Destination().(*[]int)...)
	}
	flag = isf

	return
}
