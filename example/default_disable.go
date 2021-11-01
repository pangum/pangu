package main_test

import (
	`github.com/pangum/pangu`
)

func defaultDisable() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.DisableDefault(),
	).Run(newBootstrap))
}
