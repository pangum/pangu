package command

import (
	"github.com/goexl/gox"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/runtime"
)

var _ = New

// Default 命令基础
// 可以通过匿名继承的方式快速创建命令
type Default struct {
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
		base: &Default{
			name: name,
		},
	}
}

func (b *Default) Name() string {
	return b.name
}

func (b *Default) Aliases() []string {
	return b.aliases
}

func (b *Default) Usage() string {
	return b.usage
}

func (b *Default) Run(_ *runtime.Context) (err error) {
	return
}

func (b *Default) Arguments() (args app.Arguments) {
	return
}

func (b *Default) Subcommands() (commands app.Commands) {
	return
}

func (b *Default) Description() string {
	return b.description
}

func (b *Default) Category() string {
	return b.category
}

func (b *Default) Hidden() bool {
	return b.hidden
}
