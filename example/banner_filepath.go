package main_test

import (
	`github.com/storezhang/pangu`
)

const filepathBanner = `./baner/github.png`

func bannerWithFilepath() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.Banner(filepathBanner, pangu.BannerTypeFilepath),
	).Run(newBootstrap))
}
