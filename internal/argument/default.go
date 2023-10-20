package argument

import (
	"time"

	"github.com/goexl/gox"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/argument/internal"
	"github.com/pangum/pangu/internal/runtime"
	"github.com/urfave/cli/v2"
)

var _ app.Argument = (*Default[int])(nil)

type Default[T Type] struct {
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
	environments []string
	// 是否是必须
	required bool
	// 是否隐藏
	hidden bool
	// 默认显示字符串
	text string
	// 动作
	action internal.Action[T]

	_ gox.CannotCopy
}

func NewDefault[T Type](name string, target *T) *Default[T] {
	return &Default[T]{
		name:         name,
		target:       target,
		environments: make([]string, 0),
	}
}

func (d *Default[T]) Name() string {
	return d.name
}

func (d *Default[T]) Default() any {
	return d.value
}

func (d *Default[T]) Target() any {
	return d.target
}

func (d *Default[T]) Aliases() []string {
	return d.aliases
}

func (d *Default[T]) Usage() string {
	return d.usage
}

func (d *Default[T]) Text() string {
	return d.text
}

func (d *Default[T]) Required() bool {
	return d.required
}

func (d *Default[T]) Hidden() bool {
	return d.hidden
}

func (d *Default[T]) Flag() (flag app.Flag) {
	switch d.Target().(type) {
	case *int:
		flag = d.int()
	case *[]int:
		flag = d.intSlice()
	case *int64:
		flag = d.int64()
	case *[]int64:
		flag = d.int64Slice()
	case *uint:
		flag = d.uint()
	case *[]uint:
		flag = d.uintSlice()
	case *uint64:
		flag = d.uint64()
	case *[]uint64:
		flag = d.uint64Slice()
	case *float64:
		flag = d.float64()
	case *[]float64:
		flag = d.float64Slice()
	case *time.Duration:
		flag = d.duration()
	case *bool:
		flag = d.bool()
	case *string:
		flag = d.string()
	case *[]string:
		flag = d.stringSlice()
	}

	return
}

func (d *Default[T]) runAction(context *cli.Context) error {
	return gox.If(nil != d.action, d.action(runtime.NewContext(context), *d.target))
}
