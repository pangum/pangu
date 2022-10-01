package arg

import (
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewBool
	_ app.Arg = (*boolArg)(nil)
)

type boolArg struct {
	*Argument

	destination *bool
}

// NewBool 创建布尔参数
func NewBool(name string, destination *bool, opts ...option) *boolArg {
	return &boolArg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (b *boolArg) Flag() (flag app.Flag) {
	bf := &cli.BoolFlag{
		Name:        b.Name(),
		Aliases:     b.Aliases(),
		Usage:       b.Usage(),
		DefaultText: b.DefaultText(),
		Required:    b.Required(),
		Hidden:      b.Hidden(),
	}
	if nil != b.Default() {
		bf.Value = b.Default().(bool)
	}
	if nil != b.Destination() {
		bf.Destination = b.Destination().(*bool)
	}
	flag = bf

	return
}
