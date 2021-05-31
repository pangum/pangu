package command

import (
	`fmt`

	`github.com/storezhang/pangu/app`
)

var _ app.Command = (*test)(nil)

type test struct{}

func newTest() test {
	return test{}
}

func (t *test) Aliases() []string {
	return []string{"t"}
}

func (t *test) Name() string {
	return "test"
}

func (t *test) Usage() string {
	return "测试命令"
}

func (t *test) Run(_ *app.Context) (err error) {
	fmt.Println("Test Command")

	return
}

func (t *test) SubCommands() (commands []app.Command) {
	return
}
