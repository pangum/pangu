package asset

import (
	_ "embed"
)

//go:embed template/help_app.tmpl
var HelpApp string

//go:embed template/help_command.tmpl
var HelpCommand string

//go:embed template/help_subcommand.tmpl
var HelpSubcommand string
