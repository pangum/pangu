package pangu

var _ option = (*optionAuthors)(nil)

type optionAuthors struct {
	authors []Author
}

// Authors 配置作者
func Authors(authors ...Author) *optionAuthors {
	return &optionAuthors{
		authors: authors,
	}
}

func (a *optionAuthors) apply(options *options) {
	options.authors = a.authors
}
