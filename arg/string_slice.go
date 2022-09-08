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
	*Argument
}

// NewStrings 创建一个字符串数组参数
func NewStrings(name string, opts ...option) *stringsArg {
	return &stringsArg{
		Argument: New(name, opts...),
	}
}

func (s *stringsArg) Flag() app.Flag {
	return &cli.StringSliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Destination: cli.NewStringSlice(*s.Destination().(*[]string)...),
		Value:       cli.NewStringSlice(s.Default().([]string)...),
		DefaultText: s.DefaultText(),
		Required:    s.Required(),
		Hidden:      s.Hidden(),
	}
}
