package main

import (
	"github.com/urfave/cli/v2"
	"iris_cli/commands"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "iris_cli"
	app.Usage = "工具集"
	app.Version = commands.GetVersion()
	app.Authors = []*cli.Author{{
		Name:  "xdm",
		Email: "2399364196@qq.com",
	}}
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "查看帮助",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "GoCMD Version",
	}
	app.Commands = commands.InitCommands()
	err := app.Run(os.Args)
	if err != nil {
		log.Print(err)
	}
}
