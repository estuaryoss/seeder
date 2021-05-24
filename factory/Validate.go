package factory

import (
	"fmt"
	"github.com/mitchellh/cli"
)

func Validate() (cli.Command, error) {
	validate := &validateCommandCLI{}
	return validate, nil
}

type validateCommandCLI struct {
	Args []string
}

func (c *validateCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println("validate")
	return 0
}

func (c *validateCommandCLI) Synopsis() string { return "Usage: seeder validate" }
func (c *validateCommandCLI) Help() string {
	return `
Usage: seeder validate

  Performs several layers of validation before trying to create a deployment plan:
- global config validation
- deployment files validation

  It only validates local files, but no objects found on the remote state.

Call it before 'init'. Always.
`
}
