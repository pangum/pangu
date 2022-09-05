package banner

import (
	_ "embed"

	"github.com/pangum/pangu"
	"github.com/pangum/pangu/example/bootstrap"
)

//go:embed github.png
var binaryBanner []byte
var _ = binary

func binary() {
	panic(pangu.New(
		pangu.Banner(binaryBanner, pangu.BannerTypeBinary),
	).Run(bootstrap.NewBootstrap))
}
