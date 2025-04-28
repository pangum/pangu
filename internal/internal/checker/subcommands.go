package checker

import (
	"github.com/harluo/boot/internal/application"
)

type Subcommands interface {
	Commands() []application.Command
}
