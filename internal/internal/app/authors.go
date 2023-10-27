package app

import (
	"github.com/urfave/cli/v2"
)

type Authors []*Author

func (as *Authors) Cli() (authors []*cli.Author) {
	authors = make([]*cli.Author, 0, len(*as))
	for _, a := range *as {
		author := new(cli.Author)
		author.Name = a.Name
		author.Email = a.Email
		authors = append(authors, author)
	}

	return
}
