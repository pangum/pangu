package command

import (
	`fmt`

	`github.com/storezhang/pangu/app`
)

var _ app.Command = (*Test)(nil)

type Test struct{}

func newTest() *Test {
	return &Test{}
}

func (t *Test) Aliases() []string {
	return []string{"t"}
}

func (t *Test) Name() string {
	return "Test"
}

func (t *Test) Usage() string {
	return "测试命令"
}

func (t *Test) Run(_ *app.Context) (err error) {
	fmt.Println("Test Command")

	return
}

func (t *Test) SubCommands() (commands []app.Command) {
	return
}
