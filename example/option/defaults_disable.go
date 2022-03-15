package option

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/bootstrap`
)

var _ = defaultsDisable

func defaultsDisable() {
	panic(pangu.New(
		pangu.DisableDefaults(),
		// pangu.DisableDefault(),
	).Run(bootstrap.NewBootstrap))
}
