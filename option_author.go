package pangu

var (
	_        = Author
	_ option = (*optionAuthor)(nil)
)

type optionAuthor struct {
	name  string
	email string
}

// Author 作者
func Author(name string, email string) *optionAuthor {
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
