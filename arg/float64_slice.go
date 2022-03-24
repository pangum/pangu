package arg

import (
	`github.com/pangum/pangu/app`
	`github.com/urfave/cli/v2`
)

var (
	_         = NewFloat64s
	_ app.Arg = (*float64sArg)(nil)
)

type float64sArg struct {
	*base

	// 值
	values []float64
}

// NewFloat64s 创建一个浮点数组参数
func NewFloat64s(base *base, values ...float64) *float64sArg {
	return &float64sArg{
		base:   base,
		values: values,
	}
}

func (s *float64sArg) Value() interface{} {
	return s.values
}

func (s *float64sArg) Flag() app.Flag {
	return &cli.Float64SliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Value:       cli.NewFloat64Slice(s.Value().([]float64)...),
		DefaultText: s.DefaultText(),
	}
}
