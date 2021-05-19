package factory

import (
	"fmt"
	"github.com/mitchellh/cli"
)

func Destroy() (cli.Command, error) {
	destroy := &destroyCommandCLI{}
	return destroy, nil
}

type destroyCommandCLI struct {
	Args []string
}

func (c *destroyCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println("destroy")
	return 0
}

func (c *destroyCommandCLI) Synopsis() string { return "" }
func (c *destroyCommandCLI) Help() string     { return "" }
