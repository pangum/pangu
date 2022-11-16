package arg

import (
	"time"

	"github.com/goexl/gox"
	"github.com/pangum/pangu/app"
	"github.com/urfave/cli/v2"
)

var _ app.Argument = (*argument[int])(nil)

type argument[T argumentType] struct {
	// 名称
	name string
	// 默认值
	value T
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

func (a *argument[T]) Target() (target any) {
	return
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

func (a *argument[T]) Flag() (flag app.Flag) {
	target := (any)(a.target)
	value := (any)(a.value)

	switch target.(type) {
	case *int:
		flag = a.int(target, value)
	case *[]int:
		flag = a.intSlice(target, value)
	case *int64:
		flag = a.int64(target, value)
	case *[]int64:
		flag = a.int64Slice(target, value)
	case *float64:
		flag = a.float64(target, value)
	case *time.Duration:
		flag = a.duration(target, value)
	case *bool:
		flag = a.bool(target, value)
	case *string:
		flag = a.string(target, value)
	case *[]string:
		flag = a.stringSlice(target, value)
	}

	return
}

func (a *argument[T]) runAction(ctx *cli.Context) (err error) {
	if nil == a.action {
		return
	}
	err = a.action(app.NewContext(ctx), *a.target)

	return
}
