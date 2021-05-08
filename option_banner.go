package pangu

var _ option = (*optionBanner)(nil)

type optionBanner struct {
	data       interface{}
	bannerType BannerType
}

// Banner 配置标志
func Banner(data interface{}, bannerType BannerType) *optionBanner {
	return &optionBanner{
		data:       data,
		bannerType: bannerType,
	}
}

func (b *optionBanner) apply(options *options) {
	options.banner = banner{
		data:       b.data,
		bannerType: b.bannerType,
	}
}
