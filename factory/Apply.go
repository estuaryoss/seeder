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

func (c *applyCommandCLI) Synopsis() string { return "Usage: seeder apply" }
func (c *applyCommandCLI) Help() string {
	return `
Usage: seeder apply

    Applies the deployment plan. It will bring up all your deployments and sync the current plan with the remote state.

Quote: 'You don't go to war without a plan'

Call it after 'plan', before configuration is applied on remote. Always.`
}
