package main_test

import (
	`embed`
	_ `embed`

	`github.com/storezhang/pangu`
)

//go:embed banner/txt.txt
var fileBanner embed.FS

func bannerWithFile() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.Banner(fileBanner, pangu.BannerTypeFile),
	).Run(newBootstrap))
}
