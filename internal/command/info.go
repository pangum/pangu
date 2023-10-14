package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/pangum/pangu/internal/app"
	"github.com/pangum/pangu/internal/runtime"
)

var _ app.Command = (*Info)(nil)

// Info 描述一个打印版本信息的命令
type Info struct {
	*Base
}

func newInfo() *Info {
	return &Info{
		Base: New("info").Usage(`打印应用程序信息`).Aliases(`i`, `information`).Build(),
	}
}

func (i *Info) Run(_ *app.Context) (err error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`%s\n`, strings.Repeat(`-`, 120)))
	sb.WriteString(fmt.Sprintf(`Name: %s\n`, runtime.Name))
	sb.WriteString(fmt.Sprintf(`Version: %s\n`, runtime.Version))
	sb.WriteString(fmt.Sprintf(`Build: %s\n`, runtime.Build))
	sb.WriteString(fmt.Sprintf(`Timestamp: %s\n`, runtime.Timestamp))
	sb.WriteString(fmt.Sprintf(`Revision: %s\n`, runtime.Revision))
	sb.WriteString(fmt.Sprintf(`Branch: %s\n`, runtime.Branch))
	sb.WriteString(fmt.Sprintf(`Golang: %s\n`, runtime.Golang))
	sb.WriteString(fmt.Sprintf(`%s\n`, strings.Repeat(`-`, 120)))

	fmt.Print(sb.String())
	// 刷新缓存，保证以上信息是一起被输出
	os.Stdout.Sync()

	return
}
