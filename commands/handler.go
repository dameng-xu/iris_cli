package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path/filepath"
)

func NewAction() cli.ActionFunc {
	return func(ctx *cli.Context) (err error) {
		f.Name = ctx.Args().First()
		fmt.Println(f.Name)
		fmt.Println(f.Path)
		fmt.Println(f.Host)
		fmt.Println(f.LogPath)
		if f.Path != "" {
			if f.Path, err = filepath.Abs(f.Path); err != nil {
				return
			}
			f.Path = filepath.Join(f.Path, f.Name)
		} else {
			pwd, _ := os.Getwd()
			f.Path = filepath.Join(pwd, f.Name)
		}
		f.ModPrefix = modPath(f.Path)
		// creata a project
		if err := create(); err != nil {
			return err
		}

		for _, t := range toolList() {
			if !t.installed() || t.needUpdated() {
				t.install(f.Path)
			}
		}

		fmt.Printf("Project: %s\n", f.Name)
		fmt.Printf("Directory: %s\n\n", f.Path)
		fmt.Println("项目创建成功.")
		return nil
	}
}

func BuildAction(c *cli.Context) error {
	base, err := os.Getwd()
	if err != nil {
		return err
	}
	args := append([]string{"build", "-i"}, c.Args().Slice()...)
	cmd := exec.Command("go", args...)
	cmd.Dir = buildDir(base, "cmd", 5)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("directory: %s\n", cmd.Dir)
	fmt.Printf("iris: %s\n", Version)
	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("build success.")
	return nil
}

func RunAction(c *cli.Context) error {
	base, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := buildDir(base, "cmd", 5)
	args := append([]string{"run", "main.go"}, c.Args().Slice()...)
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
