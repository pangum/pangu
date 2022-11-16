package arg

import (
	"github.com/urfave/cli/v2"
)

func (a *argument[T]) float64(target any, value any) (flag *cli.Float64Flag) {
	flag = new(cli.Float64Flag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = value.(float64)

	_target := target.(*float64)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values float64) error {
		return a.runAction(ctx)
	}

	return
}
