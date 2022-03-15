package option

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/bootstrap`
)

var _ = copyright

func copyright() {
	panic(pangu.New(
		pangu.Copyright(`https://pangu.archtech.studio`),
	).Run(bootstrap.NewBootstrap))
}
