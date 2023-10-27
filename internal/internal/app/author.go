package app

type Author struct {
	Name  string
	Email string
}

func NewAuthor(name string, email string) *Author {
	return &Author{
		Name:  name,
		Email: email,
	}
}
