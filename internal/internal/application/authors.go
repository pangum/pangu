package application

import (
	"github.com/urfave/cli/v2"
)

type Authors []*Author

func (a *Authors) Cli() (authors []*cli.Author) {
	authors = make([]*cli.Author, 0, len(*a))
	for _, a := range *a {
		author := new(cli.Author)
		author.Name = a.Name
		author.Email = a.Email
		authors = append(authors, author)
	}

	return
}
