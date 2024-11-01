package app

import (
	"github.com/urfave/cli/v2"
)

// Arguments 参数列表
type Arguments []Argument

func (a Arguments) Add(required Argument, optionals ...Argument) (arguments Arguments) {
	arguments = make([]Argument, len(a)+len(optionals))
	index := 0
	for _, argument := range append(Arguments{required}, optionals...) {
		arguments[index] = argument
		index++
	}
	for _, arg := range a {
		arguments[index] = arg
		index++
	}

	return
}

func (a Arguments) Append(required Argument, optionals ...Argument) Arguments {
	return append(a, append(Arguments{required}, optionals...)...)
}

func (a Arguments) Flags() (flags []cli.Flag) {
	flags = make([]cli.Flag, 0, len(a))
	for _, argument := range a {
		flags = append(flags, argument.Flag())
	}

	return
}
