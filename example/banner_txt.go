package main_test

import (
	_ `embed`

	`github.com/pangum/pangu`
)

//go:embed banner/txt.txt
var txtBanner string

func bannerWithTxt() {
	panic(pangu.New(
		pangu.Name("example"),
		pangu.Banner(txtBanner, pangu.BannerTypeTxt),
	).Run(newBootstrap))
}
