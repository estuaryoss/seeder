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

    Applies a deployment plan. It will bring up all your deployments and sync the local plan with the remote state.
It will run 'plan' action before, and prompt the user to accept it.

Best practice is to call it after 'plan', before configuration is applied on remote.`
}
