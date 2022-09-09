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

	destination *string
}

// NewString 创建一个字符串参数
func NewString(name string, destination *string, opts ...option) *stringArg {
	return &stringArg{
		Argument: New(name, opts...),
	}
}

func (s *stringArg) Destination() any {
	return s.destination
}

func (s *stringArg) Flag() (flag app.Flag) {
	sf := &cli.StringFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		DefaultText: s.DefaultText(),
		Required:    s.Required(),
		Hidden:      s.Hidden(),
	}
	if nil != s.Default() {
		sf.Value = s.Default().(string)
	}
	if nil != s.Destination() {
		sf.Destination = s.Destination().(*string)
	}
	flag = sf

	return
}
