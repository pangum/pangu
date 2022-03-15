package banner

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/bootstrap`
)

const filepathBanner = `./baner/github.png`

var _ = filepath

func filepath() {
	panic(pangu.New(
		pangu.Banner(filepathBanner, pangu.BannerTypeFilepath),
	).Run(bootstrap.NewBootstrap))
}
