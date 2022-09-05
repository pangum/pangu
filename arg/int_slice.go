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

func (s *intsArg) Value() interface{} {
	return s.values
}

func (s *intsArg) Flag() app.Flag {
	return &cli.IntSliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Value:       cli.NewIntSlice(s.Value().([]int)...),
		DefaultText: s.DefaultText(),
	}
}
