package banner

import (
	"embed"
	_ "embed"

	"github.com/pangum/pangu"
	"github.com/pangum/pangu/example/bootstrap"
)

//go:embed txt.txt
var fileBanner embed.FS
var _ = file

func file() {
	panic(pangu.New(
		pangu.Banner(fileBanner, pangu.BannerTypeFile),
	).Run(bootstrap.NewBootstrap))
}
