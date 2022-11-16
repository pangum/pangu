package arg

import (
	"time"

	"github.com/urfave/cli/v2"
)

func (a *argument[T]) duration() (flag *cli.DurationFlag) {
	flag = new(cli.DurationFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = a.Default().(time.Duration)

	_target := a.Target().(*time.Duration)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values time.Duration) error {
		return a.runAction(ctx)
	}

	return
}
