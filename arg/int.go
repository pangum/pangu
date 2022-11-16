package arg

import (
	"github.com/urfave/cli/v2"
)

func (a *argument[T]) int() (flag *cli.IntFlag) {
	flag = new(cli.IntFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = a.Default().(int)

	_target := a.Target().(*int)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values int) error {
		return a.runAction(ctx)
	}

	return
}

func (a *argument[T]) intSlice() (flag *cli.IntSliceFlag) {
	flag = new(cli.IntSliceFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs

	defaults := a.Default().([]int)
	flag.Value = cli.NewIntSlice(defaults...)
	flag.Action = func(ctx *cli.Context, values []int) (err error) {
		_target := a.Target().(*[]int)
		*_target = values
		err = a.runAction(ctx)
		*_target = append(*_target, defaults...)

		return
	}

	return
}

func (a *argument[T]) int64() (flag *cli.Int64Flag) {
	flag = new(cli.Int64Flag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = a.Default().(int64)

	_target := a.Target().(*int64)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values int64) error {
		return a.runAction(ctx)
	}

	return
}

func (a *argument[T]) int64Slice() (flag *cli.Int64SliceFlag) {
	flag = new(cli.Int64SliceFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs

	defaults := a.Default().([]int64)
	flag.Value = cli.NewInt64Slice(defaults...)
	flag.Action = func(ctx *cli.Context, values []int64) (err error) {
		_target := a.Target().(*[]int64)
		*_target = values
		err = a.runAction(ctx)
		*_target = append(*_target, defaults...)

		return
	}

	return
}
