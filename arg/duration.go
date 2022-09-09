package arg

import (
	"time"

	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var (
	_         = NewDuration
	_ app.Arg = (*durationArg)(nil)
)

type durationArg struct {
	*Argument

	destination *time.Duration
}

// NewDuration 创建时间区间参数
func NewDuration(name string, destination *time.Duration, opts ...option) *durationArg {
	return &durationArg{
		Argument:    New(name, opts...),
		destination: destination,
	}
}

func (d *durationArg) Destination() any {
	return d.destination
}

func (d *durationArg) Flag() (flag app.Flag) {
	ff := &cli.DurationFlag{
		Name:        d.Name(),
		Aliases:     d.Aliases(),
		Usage:       d.Usage(),
		DefaultText: d.DefaultText(),
		Required:    d.Required(),
		Hidden:      d.Hidden(),
	}
	if nil != d.Default() {
		ff.Value = d.Default().(time.Duration)
	}
	if nil != d.Destination() {
		ff.Destination = d.Destination().(*time.Duration)
	}
	flag = ff

	return
}
