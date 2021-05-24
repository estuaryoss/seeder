package factory

import (
	"fmt"
	"github.com/mitchellh/cli"
)

func Show() (cli.Command, error) {
	show := &showCommandCLI{}
	return show, nil
}

type showCommandCLI struct {
	Args []string
}

func (c *showCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println("show")
	return 0
}

func (c *showCommandCLI) Synopsis() string { return "Usage: seeder show" }
func (c *showCommandCLI) Help() string {
	return `
Usage: seeder show

    Gives full information about the current remote state. It also warns the user about possible leeway 
between the latest locally user plan and the remote state.
	
	If there is a leeway between plan and state, the user can review the deployment file, debug the deployment,
or simply recreate the affected deployment with a 'plan' and 'apply' commands.     
`
}
