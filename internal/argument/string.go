package argument

import (
	"github.com/urfave/cli/v2"
)

func (d *Default[T]) string() (flag *cli.StringFlag) {
	flag = new(cli.StringFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments
	flag.Value = d.Default().(string)

	_target := d.Target().(*string)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values string) error {
		return d.runAction(ctx)
	}

	return
}

func (d *Default[T]) stringSlice() (flag *cli.StringSliceFlag) {
	flag = new(cli.StringSliceFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments

	defaults := d.Default().([]string)
	flag.Value = cli.NewStringSlice(defaults...)
	flag.Action = func(ctx *cli.Context, values []string) (err error) {
		target := d.Target().(*[]string)
		*target = values
		err = d.runAction(ctx)
		if d.addable {
			*target = append(*target, defaults...)
		}

		return
	}

	return
}
