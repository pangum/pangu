package param

import (
	"github.com/pangum/pangu/internal/constant"
)

type Tag struct {
	Default string
}

func NewTag() *Tag {
	return &Tag{
		Default: constant.DefaultTag,
	}
}