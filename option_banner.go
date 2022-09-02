package pangu

var (
	_ option = (*optionBanner)(nil)
)

type optionBanner struct {
	data interface{}
	typ  bannerType
}

// Banner 标志
func Banner(data interface{}, bannerType bannerType) *optionBanner {
	return &optionBanner{
		data: data,
		typ:  bannerType,
	}
}

func (b *optionBanner) apply(options *options) {
	options.banner = _banner{
		data: b.data,
		typ:  b.typ,
	}
}
