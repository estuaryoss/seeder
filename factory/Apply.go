package factory

import (
	"fmt"
	"github.com/mitchellh/cli"
)

func Apply() (cli.Command, error) {
	apply := &applyCommandCLI{}
	return apply, nil
}

type applyCommandCLI struct {
	Args []string
}

func (c *applyCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println("apply")
	return 0
}

func (c *applyCommandCLI) Synopsis() string { return "Create or Update the Deployments" }
func (c *applyCommandCLI) Help() string     { return "" }
