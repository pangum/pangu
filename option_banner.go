package pangu

var _ option = (*optionBanner)(nil)

type optionBanner struct {
	content    string
	bannerType BannerType
}

// Banner 配置标志
func Banner(content string, bannerType BannerType) *optionBanner {
	return &optionBanner{
		content:    content,
		bannerType: bannerType,
	}
}

func (b *optionBanner) apply(options *options) {
	options.banner = banner{
		content:    b.content,
		bannerType: b.bannerType,
	}
}
