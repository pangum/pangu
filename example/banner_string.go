package main_test

import (
	`github.com/storezhang/pangu`
)

const stringBanner = `example`

func bannerWithString() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.Banner(stringBanner, pangu.BannerTypeString),
	).Run(newBootstrap))
}
