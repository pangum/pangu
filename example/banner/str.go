package banner

import (
	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/bootstrap`
)

const stringBanner = `example`

var _ = str

func str() {
	panic(pangu.New(
		pangu.Banner(stringBanner, pangu.BannerTypeString),
	).Run(bootstrap.NewBootstrap))
}
