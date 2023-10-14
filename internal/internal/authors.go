package internal

import (
	"github.com/urfave/cli/v2"
)

type Authors []*Author

func (a *Authors) Cli() (authors []*cli.Author) {
	authors = make([]*cli.Author, 0, len(*a))
	for _, author := range *a {
		to := new(cli.Author)
		to.Name = author.Name
		to.Email = author.Email
		authors = append(authors, to)
	}

	return
}
