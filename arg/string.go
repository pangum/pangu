package arg

import (
	"github.com/urfave/cli/v2"
)

func (a *argument[T]) string() (flag *cli.StringFlag) {
	flag = new(cli.StringFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = a.Default().(string)

	_target := a.Target().(*string)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values string) error {
		return a.runAction(ctx)
	}

	return
}

func (a *argument[T]) stringSlice() (flag *cli.StringSliceFlag) {
	flag = new(cli.StringSliceFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = cli.NewStringSlice(a.Default().([]string)...)
	flag.Action = func(ctx *cli.Context, values []string) (err error) {
		_target := a.Target().(*[]string)
		if nil != _target {
			*_target = append(*_target, values...)
		}
		err = a.runAction(ctx)

		return
	}

	return
}
