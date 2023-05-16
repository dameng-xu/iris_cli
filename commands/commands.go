package commands

import (
	"github.com/urfave/cli/v2"
)

func InitCommands() []*cli.Command {
	f = new(Flags)
	return []*cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "创建新项目",
			Flags:   f.ToNewAction(),
			Action:  NewAction(),
		},
		{
			Name:            "build",
			Usage:           "iris_cli build",
			Action:          BuildAction,
			SkipFlagParsing: true,
		},
		{
			Name:            "run",
			Usage:           "iris_cli run",
			Action:          RunAction,
			SkipFlagParsing: true,
		},
	}
}
