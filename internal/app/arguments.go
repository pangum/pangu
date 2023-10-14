package app

import (
	"github.com/urfave/cli/v2"
)

// Arguments 参数列表
type Arguments []Argument

func (a Arguments) Add(args ...Argument) (arguments Arguments) {
	arguments = make([]Argument, len(a)+len(args))
	index := 0
	for _, arg := range args {
		arguments[index] = arg
		index++
	}
	for _, arg := range a {
		arguments[index] = arg
		index++
	}

	return
}

func (a Arguments) Append(args ...Argument) Arguments {
	return append(a, args...)
}

func (a Arguments) Flags() (flags []cli.Flag) {
	flags = make([]cli.Flag, 0, len(a))
	for _, argument := range a {
		flags = append(flags, argument.Flag())
	}

	return
}
