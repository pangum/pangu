package example

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/command`
	`github.com/pangum/pangu/example/rest`
)

type componentIn struct {
	pangu.In

	// 必须是公开属性
	Rest *rest.Server
	// 必须是公开属性
	Test *command.Test
}
