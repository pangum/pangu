package main_test

import (
	`github.com/pangum/pangu`
)

const filepathBanner = `./baner/github.png`

func bannerWithFilepath() {
	panic(pangu.New(
		pangu.App("example"),
		pangu.Banner(filepathBanner, pangu.BannerTypeFilepath),
	).Run(newBootstrap))
}
