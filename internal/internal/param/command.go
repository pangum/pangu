package param

type Command struct {
	Name        string
	Aliases     []string
	Usage       string
	Description string
	Category    string
	Hidden      bool
}
