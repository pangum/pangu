package main

import (
	`github.com/storezhang/pangu`
)

func main() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.Banner("example", pangu.BannerTypeAscii),
	).Run(newBootstrap))
}
