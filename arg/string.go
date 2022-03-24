package arg

import (
	`github.com/pangum/pangu/app`
	`github.com/urfave/cli/v2`
)

var (
	_         = NewString
	_ app.Arg = (*stringArg)(nil)
)

type stringArg struct {
	*base

	// 值
	value string
}

// NewString 创建一个字符串参数
func NewString(base *base, value string) *stringArg {
	return &stringArg{
		base:  base,
		value: value,
	}
}

func (s *stringArg) Value() interface{} {
	return s.value
}

func (s *stringArg) Flag() app.Flag {
	return &cli.StringFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Value:       s.Value().(string),
		DefaultText: s.DefaultText(),
	}
}
