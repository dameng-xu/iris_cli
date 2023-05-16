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
			Usage:   "创建项目",
			Flags:   f.ToNewAction(),
			Action:  NewAction(),
		},
		{
			Name:            "build",
			Aliases:         []string{"b"},
			Usage:           "编译项目",
			Action:          BuildAction,
			SkipFlagParsing: true,
		},
		{
			Name:            "run",
			Aliases:         []string{"r"},
			Usage:           "运行项目",
			Action:          RunAction,
			SkipFlagParsing: true,
		},
	}
}
