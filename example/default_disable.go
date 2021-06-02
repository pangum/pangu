package main_test

import (
	`github.com/storezhang/pangu`
)

func defaultDisable() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.DisableDefault(),
	).Run(newBootstrap))
}
