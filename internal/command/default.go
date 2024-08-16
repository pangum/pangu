package command

import (
	"context"

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
}

func New(name string) *Builder {
	return &Builder{
		base: &Default{
			name: name,
		},
	}
}

func (d *Default) Name() string {
	return d.name
}

func (d *Default) Aliases() []string {
	return d.aliases
}

func (d *Default) Usage() string {
	return d.usage
}

func (d *Default) Run(_ *runtime.Context) (err error) {
	return
}

func (d *Default) Arguments() (args app.Arguments) {
	return
}

func (d *Default) Subcommands() (commands app.Commands) {
	return
}

func (d *Default) Description() string {
	return d.description
}

func (d *Default) Category() string {
	return d.category
}

func (d *Default) Hidden() bool {
	return d.hidden
}

func (d *Default) Before(_ context.Context) (err error) {
	return
}

func (d *Default) Stop(_ context.Context) (err error) {
	return
}

func (d *Default) After(_ context.Context) (err error) {
	return
}
