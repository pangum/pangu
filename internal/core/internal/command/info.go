package command

import (
	"context"
	"os"

	"github.com/harluo/boot/internal"
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/boot/internal/internal/command"
	"github.com/harluo/boot/internal/internal/constant"
)

var _ application.Command = (*Info)(nil)

// Info 描述一个打印版本信息的命令
type Info struct {
	*command.Default
}

func NewInfo() *Info {
	return &Info{
		Default: command.New(constant.CommandInfo).Usage(`打印应用程序信息`).Aliases(`i`, `information`).Build(),
	}
}

func (i *Info) Run(_ context.Context) (err error) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{constant.HeaderName, constant.HeaderValue})
	table.Append([]string{constant.ColumnName, internal.GetName()})
	table.Append([]string{constant.ColumnVersion, internal.GetVersion()})
	table.Append([]string{constant.ColumnBuild, internal.GetBuild()})
	table.Append([]string{constant.ColumnComplied, internal.GetCompiled()})
	table.Append([]string{constant.ColumnRevision, internal.GetRevision()})
	table.Append([]string{constant.ColumnBranch, internal.GetBranch()})
	table.Append([]string{constant.ColumnRuntime, internal.Runtime})
	table.Render()

	return
}
