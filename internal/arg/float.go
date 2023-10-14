package arg

import (
	"github.com/urfave/cli/v2"
)

func (a *argument[T]) float64() (flag *cli.Float64Flag) {
	flag = new(cli.Float64Flag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = a.Default().(float64)

	_target := a.Target().(*float64)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values float64) error {
		return a.runAction(ctx)
	}

	return
}

func (a *argument[T]) float64Slice() (flag *cli.Float64SliceFlag) {
	flag = new(cli.Float64SliceFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs

	defaults := a.Default().([]float64)
	flag.Value = cli.NewFloat64Slice(defaults...)
	flag.Action = func(ctx *cli.Context, values []float64) (err error) {
		_target := a.Target().(*[]float64)
		*_target = values
		err = a.runAction(ctx)
		if a.addable {
			*_target = append(*_target, defaults...)
		}

		return
	}

	return
}
