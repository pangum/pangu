package command

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/runtime"
)

var _ app.Command = (*Info)(nil)

// Info 描述一个打印版本信息的命令
type Info struct {
	*Default
}

func NewInfo() *Info {
	return &Info{
		Default: New(constant.CommandInfo).Usage(`打印应用程序信息`).Aliases(`i`, `information`).Build(),
	}
}

func (i *Info) Run(_ *runtime.Context) (err error) {
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
