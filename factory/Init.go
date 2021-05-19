package factory

import (
	"fmt"
	"github.com/mitchellh/cli"
)

func Init() (cli.Command, error) {
	init := &initCommandCLI{}
	return init, nil
}

type initCommandCLI struct {
	Args []string
}

func (c *initCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println("init")
	return 0
}

func (c *initCommandCLI) Synopsis() string { return "" }
func (c *initCommandCLI) Help() string     { return "" }
