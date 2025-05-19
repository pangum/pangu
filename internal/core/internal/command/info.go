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
	table.Header([]string{constant.HeaderName, constant.HeaderValue})
	if ne := table.Append([]string{constant.ColumnName, internal.GetName()}); nil != ne {
		err = ne
	} else if ve := table.Append([]string{constant.ColumnVersion, internal.GetVersion()}); nil != ve {
		err = ve
	} else if bue := table.Append([]string{constant.ColumnBuild, internal.GetBuild()}); nil != bue {
		err = bue
	} else if ce := table.Append([]string{constant.ColumnComplied, internal.GetCompiled()}); nil != ce {
		err = ce
	} else if ree := table.Append([]string{constant.ColumnRevision, internal.GetRevision()}); nil != ree {
		err = ree
	} else if bre := table.Append([]string{constant.ColumnBranch, internal.GetBranch()}); nil != bre {
		err = bre
	} else if rue := table.Append([]string{constant.ColumnRuntime, internal.Runtime}); nil != rue {
		err = rue
	} else {
		err = table.Render()
	}

	return
}
