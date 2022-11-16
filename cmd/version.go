package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/info"
	"github.com/storezhang/dig"
)

var _ app.Command = (*Version)(nil)

// Version 描述一个打印版本信息的命令
type (
	Version struct {
		*Command

		version info.Version
	}

	versionIn struct {
		dig.In

		Version info.Version
	}
)

// NewVersion 创建版本信息命令
func NewVersion(in versionIn) *Version {
	return &Version{
		Command: New("version").Usage("打印应用程序版本").Aliases("v", "ver").Build(),

		version: in.Version,
	}
}

func (v *Version) Run(_ *app.Context) (err error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`%s\n`, strings.Repeat(`-`, 120)))
	sb.WriteString(fmt.Sprintf(`Version: %s\n`, v.version))
	sb.WriteString(fmt.Sprintf(`%s\n`, strings.Repeat(`-`, 120)))

	fmt.Print(sb.String())
	// 刷新缓存，保证以上信息是一起被输出
	os.Stdout.Sync()

	return
}
