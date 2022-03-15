package pangu

var (
	_        = Authors
	_ option = (*optionAuthor)(nil)
)

type optionAuthor struct {
	name  string
	email string
}

// Authors 配置作者
func Authors(name string, email string) *optionAuthor {
	return &optionAuthor{
		name:  name,
		email: email,
	}
}

func (a *optionAuthor) apply(options *options) {
	options.authors = append(options.authors, &author{
		name:  a.name,
		email: a.email,
	})
}
