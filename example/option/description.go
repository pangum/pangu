package option

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/example/bootstrap"
)

var _ = description

func description() {
	panic(pangu.New(
		pangu.Description(`一个使用盘古框架完成的应用`),
	).Run(bootstrap.NewBootstrap))
}
