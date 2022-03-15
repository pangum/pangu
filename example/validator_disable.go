package example

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/bootstrap`
)

func validatorDisable() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.DisableValidate(),
	).Run(bootstrap.newBootstrap))
}
