package argument

import (
	"github.com/urfave/cli/v2"
)

func (d *Default[T]) int() (flag *cli.IntFlag) {
	flag = new(cli.IntFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments
	flag.Value = d.Default().(int)

	target := d.Target().(*int)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, values int) error {
		return d.runAction(ctx)
	}

	return
}

func (d *Default[T]) intSlice() (flag *cli.IntSliceFlag) {
	flag = new(cli.IntSliceFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments

	defaults := d.Default().([]int)
	flag.Value = cli.NewIntSlice(defaults...)
	flag.Action = func(ctx *cli.Context, values []int) (err error) {
		target := d.Target().(*[]int)
		*target = values
		err = d.runAction(ctx)
		if d.addable {
			*target = append(*target, defaults...)
		}

		return
	}

	return
}

func (d *Default[T]) int64() (flag *cli.Int64Flag) {
	flag = new(cli.Int64Flag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments
	flag.Value = d.Default().(int64)

	_target := d.Target().(*int64)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values int64) error {
		return d.runAction(ctx)
	}

	return
}

func (d *Default[T]) int64Slice() (flag *cli.Int64SliceFlag) {
	flag = new(cli.Int64SliceFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments

	defaults := d.Default().([]int64)
	flag.Value = cli.NewInt64Slice(defaults...)
	flag.Action = func(ctx *cli.Context, values []int64) (err error) {
		_target := d.Target().(*[]int64)
		*_target = values
		err = d.runAction(ctx)
		if d.addable {
			*_target = append(*_target, defaults...)
		}

		return
	}

	return
}
