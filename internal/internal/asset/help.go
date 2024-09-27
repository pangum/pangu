package asset

import (
	_ "embed"
)

//go:embed template/help_app.gohtml
var HelpApp string

//go:embed template/help_command.gohtml
var HelpCommand string

//go:embed template/help_subcommand.gohtml
var HelpSubcommand string
