package checker

import (
	"github.com/harluo/boot/internal/application"
)

type Commands interface {
	Commands() []application.Command
}
