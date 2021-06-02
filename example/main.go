package main_test

import (
	`github.com/storezhang/pangu`
	_ `github.com/storezhang/pangu/example/plugin`
)

func main() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.Banner("example", pangu.BannerTypeAscii),
		pangu.Usage("盘古框架例子使用方法"),
		pangu.Description("盘古框架例子描述"),
		pangu.Authors(pangu.Author{
			Name:  "storezhang",
			Email: "storezhang@gmail.com",
		}),
		pangu.Copyright("https://pangu.archtech.studio"),
	).Run(newBootstrap))
}
