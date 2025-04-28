package command

import (
	"context"
	"os"

	"github.com/harluo/boot/internal"
	"github.com/harluo/boot/internal/internal/constant"
	"github.com/olekukonko/tablewriter"
)

type Info struct {
	// 无字段
}

func newInfo() *Info {
	return &Info{
		// 无字段
	}
}

func (i *Info) Name() string {
	return "info"
}

func (i *Info) Usage() string {
	return `打印应用程序信息`
}

func (i *Info) Aliases() []string {
	return []string{
		`i`,
		`information`,
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
