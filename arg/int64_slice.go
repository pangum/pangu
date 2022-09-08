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
	*base

	// 值
	values []int64
}

// NewInt64s 创建一个整形数组参数
func NewInt64s(base *base, values ...int64) *int64sArg {
	return &int64sArg{
		base:   base,
		values: values,
	}
}

func (i *int64sArg) Default() interface{} {
	return i.values
}

func (i *int64sArg) Flag() app.Flag {
	return &cli.Int64SliceFlag{
		Name:        i.Name(),
		Aliases:     i.Aliases(),
		Usage:       i.Usage(),
		Destination: cli.NewInt64Slice(*i.Destination().(*[]int64)...),
		Value:       cli.NewInt64Slice(i.Default().([]int64)...),
		DefaultText: i.DefaultText(),
		Required:    i.Required(),
		Hidden:      i.Hidden(),
	}
}
