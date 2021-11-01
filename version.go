package pangu

import (
	`runtime`

	`github.com/pangum/pangu/info`
)

var (
	// AppName 应用名称
	AppName = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.AppName=$APP_NAME"来注入值`
	// AppVersion 应用版本
	AppVersion = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.AppVersion=$APP_VERSION"来注入值`
	// BuildVersion 编译版本
	BuildVersion = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.BuildVersion=$BUILD_VERSION"来注入值`
	// BuildTime 编译时间
	BuildTime = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.BuildTime=$BUILD_TIME"来注入值`
	// ScmRevision 分支版本
	ScmRevision = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.ScmRevision=$SCM_REVISION"来注入值`
	// ScmBranch 分支名称
	ScmBranch = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.ScmBranch=$SCM_BRANCH"来注入值`
	// GoVersion Go版本信息
	GoVersion = runtime.Version()
)

func appName() info.AppName {
	return info.AppName(AppName)
}

func appVersion() info.AppVersion {
	return info.AppVersion(AppVersion)
}

func buildVersion() info.BuildVersion {
	return info.BuildVersion(BuildVersion)
}

func buildTime() info.BuildTime {
	return info.BuildTime(BuildTime)
}

func scmRevision() info.ScmRevision {
	return info.ScmRevision(ScmRevision)
}

func scmBranch() info.ScmBranch {
	return info.ScmBranch(ScmBranch)
}

func goVersion() info.GoVersion {
	return info.GoVersion(GoVersion)
}
