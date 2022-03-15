package main_test

import (
	_ `embed`

	`github.com/pangum/pangu`
)

//go:embed banner/github.png
var binaryBanner []byte

func bannerWithBinary() {
	panic(pangu.New(
		pangu.App("example"),
		pangu.Banner(binaryBanner, pangu.BannerTypeBinary),
	).Run(newBootstrap))
}
