package param

import (
	"github.com/harluo/boot/internal/internal/constant"
)

type Tag struct {
	Default string
}

func NewTag() *Tag {
	return &Tag{
		Default: constant.DefaultTag,
	}
}
