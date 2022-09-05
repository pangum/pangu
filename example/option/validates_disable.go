package option

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/example/bootstrap"
)

var _ = validatesDisable

func validatesDisable() {
	panic(pangu.New(
		pangu.DisableValidates(),
		// or
		// pangu.DisableValidate(),
	).Run(bootstrap.NewBootstrap))
}
