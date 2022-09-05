package example

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/example/bootstrap"
	_ "github.com/pangum/pangu/example/plugin"
)

var _ = main

func main() {
	panic(pangu.New(
		pangu.Named(`example`),
		pangu.Banner(`example`, pangu.BannerTypeAscii),
		pangu.Usage(`盘古框架例子使用方法`),
		pangu.Description(`盘古框架例子描述`),
		pangu.Author(`storezhang`, `storezhang@gmail.com`),
		pangu.Copyright(`https://pangu.pangum.tech`),
	).Run(bootstrap.NewBootstrap))
}
