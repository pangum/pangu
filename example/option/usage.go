package option

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/bootstrap`
)

var _ = usage

func usage() {
	panic(pangu.New(
		pangu.Usage(`https://pangu.pangum.tech`),
	).Run(bootstrap.NewBootstrap))
}
