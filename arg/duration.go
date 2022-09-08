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
	*base
}

// NewDuration 创建时间区间参数
func NewDuration(name string, opts ...option) *durationArg {
	return &durationArg{
		base: _new(name, opts...),
	}
}

func (b *durationArg) Flag() app.Flag {
	return &cli.DurationFlag{
		Name:        b.Name(),
		Aliases:     b.Aliases(),
		Usage:       b.Usage(),
		Destination: b.Destination().(*time.Duration),
		Value:       b.Default().(time.Duration),
		DefaultText: b.DefaultText(),
		Required:    b.Required(),
		Hidden:      b.Hidden(),
	}
}
