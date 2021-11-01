package arg

import (
	`github.com/pangum/pangu/app`
	`github.com/urfave/cli/v2`
)

var _ app.Arg = (*float64SliceArg)(nil)

type float64SliceArg struct {
	*base

	// 值
	value []float64
}

// NewFloat64Slice 创建一个浮点数组参数
func NewFloat64Slice(base *base, values ...float64) *float64SliceArg {
	return &float64SliceArg{
		base:  base,
		value: values,
	}
}

func (s *float64SliceArg) Value() interface{} {
	return s.value
}

func (s *float64SliceArg) ParseFlag() app.Flag {
	return &cli.Float64SliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Value:       cli.NewFloat64Slice(s.Value().([]float64)...),
		DefaultText: s.DefaultText(),
	}
}
