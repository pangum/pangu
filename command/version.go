package command

import (
	`fmt`
	`strings`

	`github.com/pangum/pangu/app`
	`github.com/pangum/pangu/info`
)

var _ app.Command = (*Version)(nil)

// Version 描述一个打印版本信息的命令
type Version struct {
	Base

	appName      info.AppName
	appVersion   info.AppVersion
	buildVersion info.BuildVersion
	buildTime    info.BuildTime
	scmRevision  info.ScmRevision
	scmBranch    info.ScmBranch
	goVersion    info.GoVersion
}

// NewVersion 创建版本信息命令
func NewVersion(
	appName info.AppName, appVersion info.AppVersion,
	buildVersion info.BuildVersion, buildTime info.BuildTime,
	scmRevision info.ScmRevision, scmBranch info.ScmBranch,
	goVersion info.GoVersion,
) *Version {
	return &Version{
		Base: Base{
			name:    "version",
			aliases: []string{"v"},
			usage:   "打印版本信息",
		},

		appName:      appName,
		appVersion:   appVersion,
		buildVersion: buildVersion,
		buildTime:    buildTime,
		scmRevision:  scmRevision,
		scmBranch:    scmBranch,
		goVersion:    goVersion,
	}
}

func (v *Version) Run(_ *app.Context) (err error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", 120)))
	sb.WriteString(fmt.Sprintf("App    name: %s\n", v.appName))
	sb.WriteString(fmt.Sprintf("App    Version: %s\n", v.appVersion))
	sb.WriteString(fmt.Sprintf("Build  Version: %s\n", v.buildVersion))
	sb.WriteString(fmt.Sprintf("Build  Time: %s\n", v.buildTime))
	sb.WriteString(fmt.Sprintf("Scm    Revision: %s\n", v.scmRevision))
	sb.WriteString(fmt.Sprintf("Scm    Branch: %s\n", v.scmBranch))
	sb.WriteString(fmt.Sprintf("Golang Version: %s\n", v.goVersion))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", 120)))

	fmt.Print(sb.String())

	return
}
