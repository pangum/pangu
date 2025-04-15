package builder

import (
	"github.com/pangum/core/internal/internal/banner"
	"github.com/pangum/core/internal/internal/param"
)

type Banner struct {
	params      *param.Banner
	application *Application
}

func newBanner(application *Application) (banner *Banner) {
	banner = new(Banner)
	banner.params = application.params.Banner
	banner.application = application

	return
}

func (b *Banner) Ascii(ascii string) (ban *Banner) {
	b.params.Data = ascii
	b.params.Type = banner.TypeAscii
	ban = b

	return
}

func (b *Banner) Binary(binary []byte) (ban *Banner) {
	b.params.Data = binary
	b.params.Type = banner.TypeBinary
	ban = b

	return
}

func (b *Banner) Build() (application *Application) {
	b.application.params.Banner = b.params
	application = b.application

	return
}
