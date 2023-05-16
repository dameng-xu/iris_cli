package commands

import "github.com/urfave/cli/v2"

type Flags struct {
	Name      string
	Path      string
	ModPrefix string
	LogPath   string
	Host      string
	Port      string
}

var f *Flags

func (f *Flags) ToNewAction() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "d",
			Value:       "",
			Usage:       "指定项目所在目录",
			Destination: &f.Path,
		},
		&cli.StringFlag{
			Name:        "log",
			Value:       "/home/moresec/logs",
			Usage:       "日志目录",
			Destination: &f.LogPath,
		},
		&cli.StringFlag{
			Name:        "host",
			Value:       "127.0.0.1",
			Usage:       "服务主机",
			Destination: &f.Host,
		},
		&cli.StringFlag{
			Name:        "port",
			Value:       "8000",
			Usage:       "服务端口",
			Destination: &f.Port,
		},
	}
}
