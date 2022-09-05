package command

import (
	"fmt"
	"strings"

	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/info"
	"github.com/storezhang/dig"
)

var _ app.Command = (*Info)(nil)

// Info 描述一个打印版本信息的命令
type (
	Info struct {
		Base

		app       info.Name
		version   info.Version
		build     info.Build
		timestamp info.Timestamp
		revision  info.Revision
		branch    info.Branch
		golang    info.Golang
	}

	infoIn struct {
		dig.In

		App       info.Name
		Version   info.Version
		Build     info.Build
		Timestamp info.Timestamp
		Revision  info.Revision
		Branch    info.Branch
		Golang    info.Golang
	}
)

// NewInfo 创建版本信息命令
func NewInfo(in infoIn) *Info {
	return &Info{
		Base: Base{
			name:    `info`,
			aliases: []string{`i`},
			usage:   `打印应用程序信息`,
		},

		app:       in.App,
		version:   in.Version,
		build:     in.Build,
		timestamp: in.Timestamp,
		revision:  in.Revision,
		branch:    in.Branch,
		golang:    in.Golang,
	}
}

func (v *Info) Run(_ *app.Context) (err error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`%s\n`, strings.Repeat(`-`, 120)))
	sb.WriteString(fmt.Sprintf(`Name: %s\n`, v.app))
	sb.WriteString(fmt.Sprintf(`Version: %s\n`, v.version))
	sb.WriteString(fmt.Sprintf(`Build: %s\n`, v.build))
	sb.WriteString(fmt.Sprintf(`Timestamp: %s\n`, v.timestamp))
	sb.WriteString(fmt.Sprintf(`Revision: %s\n`, v.revision))
	sb.WriteString(fmt.Sprintf(`Branch: %s\n`, v.branch))
	sb.WriteString(fmt.Sprintf(`Golang: %s\n`, v.golang))
	sb.WriteString(fmt.Sprintf(`%s\n`, strings.Repeat(`-`, 120)))

	fmt.Print(sb.String())

	return
}
