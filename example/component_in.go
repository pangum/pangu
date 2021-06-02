package main_test

import (
	`github.com/storezhang/pangu`
	`github.com/storezhang/pangu/example/command`
	`github.com/storezhang/pangu/example/rest`
)

type componentIn struct {
	pangu.In

	// 必须是公开属性
	Rest *rest.Server
	// 必须是公开属性
	Test *command.Test
}
