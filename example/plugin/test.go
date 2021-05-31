package plugin

import (
	`fmt`
)

type Test struct{}

func newTest() Test {
	return Test{}
}

func (t *Test) Test() {
	fmt.Println("测试插件")
}
