package argument

import (
	"github.com/urfave/cli/v2"
)

func (d *Default[T]) uint() (flag *cli.UintFlag) {
	flag = new(cli.UintFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments
	flag.Value = d.Default().(uint)

	target := d.Target().(*uint)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, _ uint) error {
		return d.runAction(ctx)
	}

	return
}

func (d *Default[T]) uintSlice() (flag *cli.UintSliceFlag) {
	flag = new(cli.UintSliceFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments

	defaults := d.Default().([]uint)
	flag.Value = cli.NewUintSlice(defaults...)
	flag.Action = func(ctx *cli.Context, values []uint) (err error) {
		_target := d.Target().(*[]uint)
		*_target = values
		err = d.runAction(ctx)
		if d.addable {
			*_target = append(*_target, defaults...)
		}

		return
	}

	return
}

func (d *Default[T]) uint64() (flag *cli.Uint64Flag) {
	flag = new(cli.Uint64Flag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments
	flag.Value = d.Default().(uint64)

	target := d.Target().(*uint64)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, _ uint64) error {
		return d.runAction(ctx)
	}

	return
}

func (d *Default[T]) uint64Slice() (flag *cli.Uint64SliceFlag) {
	flag = new(cli.Uint64SliceFlag)
	flag.Name = d.name
	flag.Aliases = d.aliases
	flag.Usage = d.usage
	flag.DefaultText = d.text
	flag.Required = d.required
	flag.Hidden = d.hidden
	flag.EnvVars = d.environments

	defaults := d.Default().([]uint64)
	flag.Value = cli.NewUint64Slice(defaults...)
	flag.Action = func(ctx *cli.Context, values []uint64) (err error) {
		target := d.Target().(*[]uint64)
		*target = values
		err = d.runAction(ctx)
		if d.addable {
			*target = append(*target, defaults...)
		}

		return
	}

	return
}
