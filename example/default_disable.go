package example

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/bootstrap`
)

func defaultDisable() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.DisableDefault(),
	).Run(bootstrap.newBootstrap))
}
