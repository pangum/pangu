package command

import (
	`fmt`
	`strings`

	`github.com/pangum/pangu/app`
	`github.com/pangum/pangu/info`
)

var _ app.Command = (*Info)(nil)

// Info 描述一个打印版本信息的命令
type Info struct {
	Base

	app       info.App
	version   info.Version
	build     info.Build
	timestamp info.Timestamp
	revision  info.Revision
	branch    info.Branch
	golang    info.Golang
}

// NewInfo 创建版本信息命令
func NewInfo(
	app info.App, version info.Version,
	build info.Build, timestamp info.Timestamp,
	revision info.Revision, branch info.Branch,
	golang info.Golang,
) *Info {
	return &Info{
		Base: Base{
			name:    "info",
			aliases: []string{"i"},
			usage:   "打印应用程序信息",
		},

		app:       app,
		version:   version,
		build:     build,
		timestamp: timestamp,
		revision:  revision,
		branch:    branch,
		golang:    golang,
	}
}

func (v *Info) Run(_ *app.Context) (err error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", 120)))
	sb.WriteString(fmt.Sprintf("App: %s\n", v.app))
	sb.WriteString(fmt.Sprintf("Version: %s\n", v.version))
	sb.WriteString(fmt.Sprintf("Build: %s\n", v.build))
	sb.WriteString(fmt.Sprintf("Timestamp: %s\n", v.timestamp))
	sb.WriteString(fmt.Sprintf("Revision: %s\n", v.revision))
	sb.WriteString(fmt.Sprintf("Branch: %s\n", v.branch))
	sb.WriteString(fmt.Sprintf("Golang: %s\n", v.golang))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", 120)))

	fmt.Print(sb.String())

	return
}
