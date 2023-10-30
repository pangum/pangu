package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/pangum/pangu/internal"
	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/runtime"
)

var _ app.Command = (*Version)(nil)

// Version 描述一个打印版本信息的命令
type Version struct {
	*Default
}

func NewVersion() *Version {
	return &Version{
		Default: New("version").Usage("打印应用程序版本").Aliases("v", "ver").Build(),
	}
}

func (v *Version) Run(_ *runtime.Context) (err error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s\n", strings.Repeat(`-`, 120)))
	sb.WriteString(fmt.Sprintf("Version: %s\n", internal.Version))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Repeat(`-`, 120)))

	fmt.Print(sb.String())
	// 刷新缓存，保证以上信息是一起被输出
	err = os.Stdout.Sync()

	return
}
