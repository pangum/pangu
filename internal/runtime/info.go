package runtime

import (
	"runtime"

	"github.com/pangum/pangu/internal/constant"
)

var (
	Name      = constant.ApplicationDefaultName
	Version   = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Version=$VERSION"来注入值`
	Build     = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Build=$BUILD"来注入值`
	Timestamp = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Timestamp=$TIMESTAMP"来注入值`
	Revision  = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Revision=$REVISION"来注入值`
	Branch    = `没有设置，请使用-ldflags "-s -X 'github.com/pangum/pangu.Branch=$BRANCH"来注入值`
	Golang    = runtime.Version()
)
