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
)

var _ application.Command = (*Version)(nil)

// Version 描述一个打印版本信息的命令
type Version struct {
	*command.Default
}

func NewVersion() *Version {
	return &Version{
		Default: command.New(constant.CommandVersion).Usage("打印应用程序版本").Aliases("v", "ver").Build(),
	}
}

func (v *Version) Run(_ context.Context) (err error) {
	builder := new(strings.Builder)
	builder.WriteString(fmt.Sprintf("%s\n", strings.Repeat(`-`, 120)))
	builder.WriteString(fmt.Sprintf("Version: %s\n", internal.Version))
	builder.WriteString(fmt.Sprintf("%s\n", strings.Repeat(`-`, 120)))

	fmt.Print(builder.String())
	// 刷新缓存，保证以上信息是一起被输出
	err = os.Stdout.Sync()

	return
}
