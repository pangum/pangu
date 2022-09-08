package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewInts
	_ app.Arg = (*intsArg)(nil)
)

type intsArg struct {
	*base

	// 值
	values []int
}

// NewInts 创建一个整形数组参数
func NewInts(base *base, values ...int) *intsArg {
	return &intsArg{
		base:   base,
		values: values,
	}
}

func (i *intsArg) Default() interface{} {
	return i.values
}

func (i *intsArg) Flag() app.Flag {
	return &cli.IntSliceFlag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Destination: cli.NewIntSlice(*i.Destination().(*[]int)...),
		Value:       cli.NewIntSlice(i.Default().([]int)...),
		DefaultText: i.DefaultText(),
		Required:    i.Required(),
		Hidden:      i.Hidden(),
	}
}
