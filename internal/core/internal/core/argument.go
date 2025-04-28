package core

import (
	"time"

	"github.com/harluo/boot/internal/application"
	"github.com/harluo/boot/internal/runtime"
	"github.com/urfave/cli/v2"
)

type Argument struct {
	argument application.Argument
	typer    *Typer
}

func NewArgument(argument application.Argument) *Argument {
	return &Argument{
		argument: argument,
		typer:    NewTyper(argument),
	}
}

func (a *Argument) Flag() (flag application.Flag) {
	switch a.argument.Target().(type) {
	case *int:
		flag = a.int()
	case *[]int:
		flag = a.intSlice()
	case *int64:
		flag = a.int64()
	case *[]int64:
		flag = a.int64Slice()
	case *uint:
		flag = a.uint()
	case *[]uint:
		flag = a.uintSlice()
	case *uint64:
		flag = a.uint64()
	case *[]uint64:
		flag = a.uint64Slice()
	case *float64:
		flag = a.float64()
	case *[]float64:
		flag = a.float64Slice()
	case *time.Duration:
		flag = a.duration()
	case *bool:
		flag = a.bool()
	case *string:
		flag = a.string()
	case *[]string:
		flag = a.stringSlice()
	}

	return
}

func (a *Argument) bool() (flag *cli.BoolFlag) {
	flag = new(cli.BoolFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()
	a.typer.Default(func(value any) {
		flag.Value = value.(bool)
	})

	target := a.argument.Target().(*bool)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, values bool) error {
		return a.typer.Run(runtime.NewContext(ctx))
	}

	return
}

func (a *Argument) duration() (flag *cli.DurationFlag) {
	flag = new(cli.DurationFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()
	a.typer.Default(func(value any) {
		flag.Value = value.(time.Duration)
	})

	target := a.argument.Target().(*time.Duration)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, values time.Duration) error {
		return a.typer.Run(runtime.NewContext(ctx))
	}

	return
}

func (a *Argument) float64() (flag *cli.Float64Flag) {
	flag = new(cli.Float64Flag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()
	a.typer.Default(func(value any) {
		flag.Value = value.(float64)
	})

	target := a.argument.Target().(*float64)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, values float64) error {
		return a.typer.Run(runtime.NewContext(ctx))
	}

	return
}

func (a *Argument) float64Slice() (flag *cli.Float64SliceFlag) {
	flag = new(cli.Float64SliceFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()

	defaults := make([]float64, 0)
	a.typer.Default(func(value any) {
		defaults = value.([]float64)
	})
	flag.Value = cli.NewFloat64Slice(defaults...)
	flag.Action = func(ctx *cli.Context, values []float64) (err error) {
		_target := a.argument.Target().(*[]float64)
		*_target = values
		err = a.typer.Run(runtime.NewContext(ctx))
		if a.typer.Addable() {
			*_target = append(*_target, defaults...)
		}

		return
	}

	return
}

func (a *Argument) int() (flag *cli.IntFlag) {
	flag = new(cli.IntFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()
	a.typer.Default(func(value any) {
		flag.Value = value.(int)
	})

	target := a.argument.Target().(*int)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, values int) error {
		return a.typer.Run(runtime.NewContext(ctx))
	}

	return
}

func (a *Argument) intSlice() (flag *cli.IntSliceFlag) {
	flag = new(cli.IntSliceFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()

	defaults := make([]int, 0)
	a.typer.Default(func(value any) {
		defaults = value.([]int)
	})
	flag.Value = cli.NewIntSlice(defaults...)
	flag.Action = func(ctx *cli.Context, values []int) (err error) {
		target := a.argument.Target().(*[]int)
		*target = values
		err = a.typer.Run(runtime.NewContext(ctx))
		if a.typer.Addable() {
			*target = append(*target, defaults...)
		}

		return
	}

	return
}

func (a *Argument) int64() (flag *cli.Int64Flag) {
	flag = new(cli.Int64Flag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()
	a.typer.Default(func(value any) {
		flag.Value = value.(int64)
	})

	_target := a.argument.Target().(*int64)
	if nil != _target {
		flag.Destination = _target
	}
	flag.Action = func(ctx *cli.Context, values int64) error {
		return a.typer.Run(runtime.NewContext(ctx))
	}

	return
}

func (a *Argument) int64Slice() (flag *cli.Int64SliceFlag) {
	flag = new(cli.Int64SliceFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()

	defaults := make([]int64, 0)
	a.typer.Default(func(value any) {
		defaults = value.([]int64)
	})
	flag.Value = cli.NewInt64Slice(defaults...)
	flag.Action = func(ctx *cli.Context, values []int64) (err error) {
		_target := a.argument.Target().(*[]int64)
		*_target = values
		err = a.typer.Run(runtime.NewContext(ctx))
		if a.typer.Addable() {
			*_target = append(*_target, defaults...)
		}

		return
	}

	return
}

func (a *Argument) string() (flag *cli.StringFlag) {
	flag = new(cli.StringFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()
	a.typer.Default(func(value any) {
		flag.Value = value.(string)
	})

	target := a.argument.Target().(*string)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, values string) error {
		return a.typer.Run(runtime.NewContext(ctx))
	}

	return
}

func (a *Argument) stringSlice() (flag *cli.StringSliceFlag) {
	flag = new(cli.StringSliceFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()

	defaults := make([]string, 0)
	a.typer.Default(func(value any) {
		defaults = value.([]string)
	})
	flag.Value = cli.NewStringSlice(defaults...)
	flag.Action = func(ctx *cli.Context, values []string) (err error) {
		target := a.argument.Target().(*[]string)
		*target = values
		err = a.typer.Run(runtime.NewContext(ctx))
		if a.typer.Addable() {
			*target = append(*target, defaults...)
		}

		return
	}

	return
}

func (a *Argument) uint() (flag *cli.UintFlag) {
	flag = new(cli.UintFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()
	a.typer.Default(func(value any) {
		flag.Value = value.(uint)
	})

	target := a.argument.Target().(*uint)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, _ uint) error {
		return a.typer.Run(runtime.NewContext(ctx))
	}

	return
}

func (a *Argument) uintSlice() (flag *cli.UintSliceFlag) {
	flag = new(cli.UintSliceFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()

	defaults := make([]uint, 0)
	a.typer.Default(func(value any) {
		defaults = value.([]uint)
	})
	flag.Value = cli.NewUintSlice(defaults...)
	flag.Action = func(ctx *cli.Context, values []uint) (err error) {
		_target := a.argument.Target().(*[]uint)
		*_target = values
		err = a.typer.Run(runtime.NewContext(ctx))
		if a.typer.Addable() {
			*_target = append(*_target, defaults...)
		}

		return
	}

	return
}

func (a *Argument) uint64() (flag *cli.Uint64Flag) {
	flag = new(cli.Uint64Flag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()
	a.typer.Default(func(value any) {
		flag.Value = value.(uint64)
	})

	target := a.argument.Target().(*uint64)
	if nil != target {
		flag.Destination = target
	}
	flag.Action = func(ctx *cli.Context, _ uint64) error {
		return a.typer.Run(runtime.NewContext(ctx))
	}

	return
}

func (a *Argument) uint64Slice() (flag *cli.Uint64SliceFlag) {
	flag = new(cli.Uint64SliceFlag)
	flag.Name = a.argument.Name()
	flag.Aliases = a.typer.Aliases()
	flag.Usage = a.typer.Usage()
	flag.DefaultText = a.typer.Text()
	flag.Required = a.typer.Required()
	flag.Hidden = a.typer.Hidden()
	flag.EnvVars = a.typer.Environments()

	defaults := make([]uint64, 0)
	a.typer.Default(func(value any) {
		defaults = value.([]uint64)
	})
	flag.Value = cli.NewUint64Slice(defaults...)
	flag.Action = func(ctx *cli.Context, values []uint64) (err error) {
		target := a.argument.Target().(*[]uint64)
		*target = values
		err = a.typer.Run(runtime.NewContext(ctx))
		if a.typer.Addable() {
			*target = append(*target, defaults...)
		}

		return
	}

	return
}
