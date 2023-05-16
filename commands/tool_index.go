package commands

import "time"

var toolIndexs = []*Tool{
	{
		Name:      "mod",
		Alias:     "mod",
		BuildTime: time.Date(2020, 3, 31, 0, 0, 0, 0, time.Local),
		Install:   "go mod tidy",
		Platform:  []string{"darwin", "linux", "windows"},
		Author:    "google",
	},
}
