package command

import (
	`fmt`

	`github.com/pangum/pangu/app`
	`github.com/pangum/pangu/command`
	`github.com/pangum/pangu/example/plugin`
)

var _ app.Command = (*Test)(nil)

type Test struct {
	command.Base

	test plugin.Test
}

func newTest(test plugin.Test) *Test {
	return &Test{
		Base: command.NewBase("test", "测试命令", "t"),
		test: test,
	}
}

func (t *Test) Run(_ *app.Context) (err error) {
	fmt.Println("Test Command")
	t.test.Test()

	return
}
