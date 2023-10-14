package arg

import (
	"time"

	"github.com/goexl/gox"
	app2 "github.com/pangum/pangu/internal/app"
	"github.com/urfave/cli/v2"
)

var _ app2.Argument = (*argument[int])(nil)

type argument[T argumentType] struct {
	// 名称
	name string
	// 默认值
	value T
	// 默认值是否追加
	addable bool
	// 目标
	target *T
	// 别名
	aliases []string
	// 使用方法
	usage string
	// 环境变量列表
	envs []string
	// 是否是必须
	required bool
	// 是否隐藏
	hidden bool
	// 默认显示字符串
	text string
	// 动作
	action action[T]

	_ gox.CannotCopy
}

func (a *argument[T]) Name() string {
	return a.name
}

func (a *argument[T]) Default() any {
	return a.value
}

func (a *argument[T]) Target() any {
	return a.target
}

func (a *argument[T]) Aliases() []string {
	return a.aliases
}

func (a *argument[T]) Usage() string {
	return a.usage
}

func (a *argument[T]) Text() string {
	return a.text
}

func (a *argument[T]) Required() bool {
	return a.required
}

func (a *argument[T]) Hidden() bool {
	return a.hidden
}

func (a *argument[T]) Flag() (flag app2.Flag) {
	switch a.Target().(type) {
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

func (a *argument[T]) runAction(ctx *cli.Context) (err error) {
	if nil == a.action {
		return
	}
	err = a.action(app2.NewContext(ctx), *a.target)

	return
}
