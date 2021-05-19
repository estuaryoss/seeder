package factory

import (
	"fmt"
	"github.com/mitchellh/cli"
)

func Plan() (cli.Command, error) {
	plan := &planCommandCLI{}
	return plan, nil
}

type planCommandCLI struct {
	Args []string
}

func (c *planCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println("plan")
	return 0
}

func (c *planCommandCLI) Synopsis() string { return "" }
func (c *planCommandCLI) Help() string     { return "" }
