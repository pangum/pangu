package param

import (
	"github.com/pangum/pangu/internal/internal/asset"
)

type Help struct {
	// 帮助信息模板
	App string
	// 命令帮助信息模板
	Command string
	// 子命令帮助信息模板
	Subcommand string
}

func newHelp() *Help {
	return &Help{
		App:        asset.HelpApp,
		Command:    asset.HelpCommand,
		Subcommand: asset.HelpSubcommand,
	}
}
