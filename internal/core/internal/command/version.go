package command

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/harluo/boot/internal"
	"github.com/harluo/boot/internal/application"
	"github.com/harluo/boot/internal/internal/command"
	"github.com/harluo/boot/internal/internal/constant"
	"github.com/harluo/boot/internal/internal/param"
)

var _ application.Command = (*Version)(nil)

// Version 描述一个打印版本信息的命令
type Version struct {
	*command.Default
}

func NewVersion() *Version {
	return &Version{
		Default: command.NewDefault(&param.Command{
			Name:  constant.CommandVersion,
			Usage: `打印应用程序版本`,
			Aliases: []string{
				`v`,
				`ver`,
			},
		}),
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
