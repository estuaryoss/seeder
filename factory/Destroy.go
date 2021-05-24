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

func (c *destroyCommandCLI) Synopsis() string { return "Usage: seeder destroy" }
func (c *destroyCommandCLI) Help() string {
	return `
Usage: seeder destroy [options] id

    Destroys the remote state and empties your local plan.

Scenarios: 
    Call it without arguments and it will erase all your remote state.
    Call it with args and destroy only the objects specified by 'id'.

`
}
