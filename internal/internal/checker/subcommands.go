package checker

import (
	"github.com/harluo/boot/internal/application"
)

type Subcommands interface {
	Subcommands() []application.Command
}
