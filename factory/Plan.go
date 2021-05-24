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

func (c *planCommandCLI) Synopsis() string { return "Usage: seeder plan" }
func (c *planCommandCLI) Help() string {
	return `
Usage: seeder plan

    Creates the deployments plan. By default, creating a plan consists of:
        -  reading remote state and comparing to the current plan
        -  proposing a set of actions in order to sync remote state with the current plan

Call it after 'init'. Always.
`
}
