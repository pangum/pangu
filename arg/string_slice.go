package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewStringSlice
	_ app.Arg = (*stringSliceArg)(nil)
)

type stringSliceArg struct {
	*Argument

	destination *[]string
}

// NewStringSlice 创建一个字符串数组参数
func NewStringSlice(name string, destination *[]string, opts ...option) *stringSliceArg {
	return &stringSliceArg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (s *stringSliceArg) Destination() any {
	return s.destination
}

func (s *stringSliceArg) Flag() (flag app.Flag) {
	sf := &cli.StringSliceFlag{
		Name:        s.Name(),
		Aliases:     s.Aliases(),
		Usage:       s.Usage(),
		DefaultText: s.DefaultText(),
		Required:    s.Required(),
		Hidden:      s.Hidden(),
	}
	if nil != s.Default() {
		sf.Value = cli.NewStringSlice((s.Default().([]string))...)
	}
	if nil != s.Destination() {
		sf.Destination = cli.NewStringSlice((*s.Destination().(*[]string))...)
	}
	flag = sf

	return
}
