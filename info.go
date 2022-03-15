package pangu

import (
	`runtime`

	`github.com/pangum/pangu/info`
)

var (
	// App 应用名称
	App = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.App=$App"来注入值`
	// Version 应用版本
	Version = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Version=$VERSION"来注入值`
	// Build 编译版本
	Build = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Build=$BUILD"来注入值`
	// Timestamp 编译时间
	Timestamp = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Timestamp=$TIMESTAMP"来注入值`
	// Revision 分支版本
	Revision = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Revision=$REVISION"来注入值`
	// Branch 分支名称
	Branch = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Branch=$BRANCH"来注入值`
	// Golang Go版本信息
	Golang = runtime.Version()
)

func app() info.App {
	return info.App(App)
}

func version() info.Version {
	return info.Version(Version)
}

func build() info.Build {
	return info.Build(Build)
}

func timestamp() info.Timestamp {
	return info.Timestamp(Timestamp)
}

func revision() info.Revision {
	return info.Revision(Revision)
}

func branch() info.Branch {
	return info.Branch(Branch)
}

func golang() info.Golang {
	return info.Golang(Golang)
}
