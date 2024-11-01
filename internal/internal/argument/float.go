package argument

import (
	"github.com/urfave/cli/v2"
)

func (d *Default[T]) float64() (flag *cli.Float64Flag) {
	flag = new(cli.Float64Flag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments
	flag.Value = d.Default().(float64)

	target := d.Target().(*float64)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, values float64) error {
		return d.runAction(ctx)
	}

	return
}

func (d *Default[T]) float64Slice() (flag *cli.Float64SliceFlag) {
	flag = new(cli.Float64SliceFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments

	defaults := d.Default().([]float64)
	flag.Value = cli.NewFloat64Slice(defaults...)
	flag.Action = func(ctx *cli.Context, values []float64) (err error) {
		_target := d.Target().(*[]float64)
		*_target = values
		err = d.runAction(ctx)
		if d.addable {
			*_target = append(*_target, defaults...)
		}

		return
	}

	return
}
