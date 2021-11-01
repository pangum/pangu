package main_test

import (
	`github.com/pangum/pangu`
)

func validatorDisable() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.DisableValidate(),
	).Run(newBootstrap))
}
