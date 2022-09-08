package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewString
	_ app.Arg = (*stringArg)(nil)
)

type stringArg struct {
	*Argument
}

// NewString 创建一个字符串参数
func NewString(name string, opts ...option) *stringArg {
	return &stringArg{
		Argument: New(name, opts...),
	}
}

func (s *stringArg) Flag() app.Flag {
	return &cli.StringFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		Destination: s.Destination().(*string),
		Value:       s.Default().(string),
		DefaultText: s.DefaultText(),
		Required:    s.Required(),
		Hidden:      s.Hidden(),
	}
}
