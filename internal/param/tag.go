package param

import (
	"github.com/pangum/pangu/internal/constant"
)

type Tag struct {
	Default string
	Set     bool
}

func NewTag() *Tag {
	return &Tag{
		Default: constant.DefaultTag,
	}
}
