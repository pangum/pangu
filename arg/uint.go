package arg

import (
	"github.com/urfave/cli/v2"
)

func (a *argument[T]) uint() (flag *cli.UintFlag) {
	flag = new(cli.UintFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = a.Default().(uint)

	_target := a.Target().(*uint)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, _ uint) error {
		return a.runAction(ctx)
	}

	return
}

func (a *argument[T]) uintSlice() (flag *cli.UintSliceFlag) {
	flag = new(cli.UintSliceFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs

	defaults := a.Default().([]uint)
	flag.Value = cli.NewUintSlice(defaults...)
	flag.Action = func(ctx *cli.Context, values []uint) (err error) {
		_target := a.Target().(*[]uint)
		*_target = values
		err = a.runAction(ctx)
		if a.addable {
			*_target = append(*_target, defaults...)
		}

		return
	}

	return
}

func (a *argument[T]) uint64() (flag *cli.Uint64Flag) {
	flag = new(cli.Uint64Flag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = a.Default().(uint64)

	_target := a.Target().(*uint64)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, _ uint64) error {
		return a.runAction(ctx)
	}

	return
}

func (a *argument[T]) uint64Slice() (flag *cli.Uint64SliceFlag) {
	flag = new(cli.Uint64SliceFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs

	defaults := a.Default().([]uint64)
	flag.Value = cli.NewUint64Slice(defaults...)
	flag.Action = func(ctx *cli.Context, values []uint64) (err error) {
		_target := a.Target().(*[]uint64)
		*_target = values
		err = a.runAction(ctx)
		if a.addable {
			*_target = append(*_target, defaults...)
		}

		return
	}

	return
}
