package argument

import (
	"github.com/urfave/cli/v2"
)

func (d *Default[T]) bool() (flag *cli.BoolFlag) {
	flag = new(cli.BoolFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments
	flag.Value = d.Default().(bool)

	target := d.Target().(*bool)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, values bool) error {
		return d.runAction(ctx)
	}

	return
}
