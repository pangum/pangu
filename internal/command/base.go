package command

import (
	"github.com/goexl/gox"
	"github.com/pangum/pangu/internal/app"
)

var _ = New

// Base 命令基础
// 可以通过匿名继承的方式快速创建命令
type Base struct {
	name        string
	aliases     []string
	usage       string
	description string
	category    string
	hidden      bool

	_ gox.CannotCopy
}

func New(name string) *Builder {
	return &Builder{
		base: &Base{
			name: name,
		},
	}
}

func (b *Base) Name() string {
	return b.name
}

func (b *Base) Aliases() []string {
	return b.aliases
}

func (b *Base) Usage() string {
	return b.usage
}

func (b *Base) Run(_ *app.Context) (err error) {
	return
}

func (b *Base) Arguments() (args app.Arguments) {
	return
}

func (b *Base) Subcommands() (commands app.Commands) {
	return
}

func (b *Base) Description() string {
	return b.description
}

func (b *Base) Category() string {
	return b.category
}

func (b *Base) Hidden() bool {
	return b.hidden
}
