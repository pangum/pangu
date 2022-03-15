package banner

import (
	_ `embed`

	`github.com/pangum/pangu`
	`github.com/pangum/pangu/example/bootstrap`
)

//go:embed txt.txt
var txtBanner string

var _ = txt

func txt() {
	panic(pangu.New(
		pangu.Banner(txtBanner, pangu.BannerTypeTxt),
	).Run(bootstrap.NewBootstrap))
}
