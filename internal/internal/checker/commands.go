package checker

import (
	"github.com/harluo/boot/internal/application"
)

type Commands interface {
	Subcommands() []application.Command
}
