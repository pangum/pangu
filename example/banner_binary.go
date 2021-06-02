package main_test

import (
	_ `embed`

	`github.com/storezhang/pangu`
)

//go:embed banner/github.png
var binaryBanner []byte

func bannerWithBinary() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.Banner(binaryBanner, pangu.BannerTypeBinary),
	).Run(newBootstrap))
}
