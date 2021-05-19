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

func (c *validateCommandCLI) Synopsis() string { return "" }
func (c *validateCommandCLI) Help() string     { return "" }
