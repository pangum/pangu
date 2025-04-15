package param

import (
	"github.com/pangum/core/internal/internal/constant"
)

type Tag struct {
	Default string
}

func NewTag() *Tag {
	return &Tag{
		Default: constant.DefaultTag,
	}
}
