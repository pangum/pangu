package arg

import (
	"github.com/urfave/cli/v2"
)

func (a *argument[T]) bool(target any, value any) (flag *cli.BoolFlag) {
	flag = new(cli.BoolFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = value.(bool)

	_target := target.(*bool)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values bool) error {
		return a.runAction(ctx)
	}

	return
}
