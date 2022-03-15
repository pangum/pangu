package option

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/bootstrap`
)

var _ = author

func author() {
	panic(pangu.New(
		pangu.Author(`storezhang`, `storezhang@gmail.com`),
	).Run(bootstrap.NewBootstrap))
}
