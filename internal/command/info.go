package command

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/command/internal/constant"
)

var _ app.Command = (*Info)(nil)

// Info 描述一个打印版本信息的命令
type Info struct {
	*Base
}

func NewInfo() *Info {
	return &Info{
		Base: New("info").Usage(`打印应用程序信息`).Aliases(`i`, `information`).Build(),
	}
}

func (i *Info) Run(_ *app.Context) (err error) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{constant.HeaderName, constant.HeaderValue})
	table.Append([]string{constant.ColumnName, internal.Name})
	table.Append([]string{constant.ColumnVersion, internal.Version})
	table.Append([]string{constant.ColumnBuild, internal.Build})
	table.Append([]string{constant.ColumnTimestamp, internal.Timestamp})
	table.Append([]string{constant.ColumnRevision, internal.Revision})
	table.Append([]string{constant.ColumnBranch, internal.Branch})
	table.Append([]string{constant.ColumnGolang, internal.Golang})
	table.Render()

	return
}
