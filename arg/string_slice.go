package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewStrings
	_ app.Arg = (*stringsArg)(nil)
)

type stringsArg struct {
	*base

	// 值
	values []string
}

// NewStrings 创建一个字符串数组参数
func NewStrings(base *base, values ...string) *stringsArg {
	return &stringsArg{
		base:   base,
		values: values,
	}
}

func (s *stringsArg) Value() interface{} {
	return s.values
}

func (s *stringsArg) Flag() app.Flag {
	return &cli.StringSliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Value:       cli.NewStringSlice(s.Value().([]string)...),
		DefaultText: s.DefaultText(),
	}
}
