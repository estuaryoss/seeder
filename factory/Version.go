package factory

import (
	"fmt"
	"github.com/mitchellh/cli"
	"seeder/constants"
)

func Version() (cli.Command, error) {
	version := &versionCommandCLI{}
	return version, nil
}

type versionCommandCLI struct {
	Args []string
}

func (c *versionCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println(constants.APP_NAME + ", version " + constants.APP_VERSION)
	return 0
}

func (c *versionCommandCLI) Synopsis() string { return "Usage: seeder version" }
func (c *versionCommandCLI) Help() string {
	return `
Usage: seeder version

  Gets the current application version of the CLI.
`
}
