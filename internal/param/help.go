package param

import (
	"github.com/pangum/pangu/internal/asset"
)

type Help struct {
	// 帮助信息模板
	App string
	// 命令帮助信息模板
	Command string
	// 子命令帮助信息模板
	Subcommand string
}

func NewHelp() *Help {
	return &Help{
		App:        asset.HelpApp,
		Command:    asset.HelpCommand,
		Subcommand: asset.HelpSubcommand,
	}
}
