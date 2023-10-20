package argument

import (
	"time"

	"github.com/urfave/cli/v2"
)

func (d *Default[T]) duration() (flag *cli.DurationFlag) {
	flag = new(cli.DurationFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments
	flag.Value = d.Default().(time.Duration)

	target := d.Target().(*time.Duration)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, values time.Duration) error {
		return d.runAction(ctx)
	}

	return
}
