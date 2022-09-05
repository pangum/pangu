package option

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/example/bootstrap"
)

var _ = named

func named() {
	panic(pangu.New(
		pangu.Named(`一个使用盘古框架完成的应用`),
	).Run(bootstrap.NewBootstrap))
}
