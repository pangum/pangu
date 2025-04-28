package command

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/harluo/boot/internal"
)

type Version struct {
	// 无字段
}

func newVersion() *Version {
	return &Version{
		// 无字段
	}
}

func (v *Version) Name() string {
	return "version"
}

func (v *Version) Usage() string {
	return `打印应用程序版本`
}

func (v *Version) Aliases() []string {
	return []string{
		"v",
		"ver",
	}
}

func (v *Version) Run(_ context.Context) (err error) {
	version := new(strings.Builder)
	version.WriteString(fmt.Sprintf("%s\n", strings.Repeat(`-`, 120)))
	version.WriteString(fmt.Sprintf("Version: %s\n", internal.Version))
	version.WriteString(fmt.Sprintf("%s\n", strings.Repeat(`-`, 120)))

	fmt.Print(version.String())
	// 刷新缓存，保证以上信息是一起被输出
	err = os.Stdout.Sync()

	return
}
