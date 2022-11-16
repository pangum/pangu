package arg

import (
	"github.com/urfave/cli/v2"
)

func (a *argument[T]) int(target any, value any) (flag *cli.IntFlag) {
	flag = new(cli.IntFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = value.(int)

	_target := target.(*int)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values int) error {
		return a.runAction(ctx)
	}

	return
}

func (a *argument[T]) intSlice(target any, value any) (flag *cli.IntSliceFlag) {
	flag = new(cli.IntSliceFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = cli.NewIntSlice(value.([]int)...)
	flag.Action = func(ctx *cli.Context, values []int) (err error) {
		_target := target.(*[]int)
		*_target = append(*_target, values...)
		err = a.runAction(ctx)

		return
	}

	return
}

func (a *argument[T]) int64(target any, value any) (flag *cli.Int64Flag) {
	flag = new(cli.Int64Flag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = value.(int64)

	_target := target.(*int64)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values int64) error {
		return a.runAction(ctx)
	}

	return
}

func (a *argument[T]) int64Slice(target any, value any) (flag *cli.Int64SliceFlag) {
	flag = new(cli.Int64SliceFlag)
	flag.Name = a.name
	flag.Aliases = a.aliases
	flag.Usage = a.usage
	flag.DefaultText = a.text
	flag.Required = a.required
	flag.Hidden = a.hidden
	flag.EnvVars = a.envs
	flag.Value = cli.NewInt64Slice(value.([]int64)...)
	flag.Action = func(ctx *cli.Context, values []int64) (err error) {
		_target := target.(*[]int64)
		*_target = append(*_target, values...)
		err = a.runAction(ctx)

		return
	}

	return
}
