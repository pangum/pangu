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

	destination *[]string
}

// NewStrings 创建一个字符串数组参数
func NewStrings(name string, destination *[]string, opts ...option) *stringsArg {
	return &stringsArg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (s *stringsArg) Destination() any {
	return s.destination
}

func (s *stringsArg) Flag() (flag app.Flag) {
	sf := &cli.StringSliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		DefaultText: s.DefaultText(),
		Required:    s.Required(),
		Hidden:      s.Hidden(),
	}
	if nil != s.Default() {
		sf.Value = cli.NewStringSlice(s.Default().([]string)...)
	}
	if nil != s.Destination() {
		sf.Destination = cli.NewStringSlice(*s.Destination().(*[]string)...)
	}
	flag = sf

	return
}
