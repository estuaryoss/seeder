package factory

import (
	"fmt"
	"github.com/mitchellh/cli"
)

func Init() (cli.Command, error) {
	init := &initCommandCLI{}
	return init, nil
}

type initCommandCLI struct {
	Args []string
}

func (c *initCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println("init")
	return 0
}

func (c *initCommandCLI) Synopsis() string { return "Usage: seeder init" }
func (c *initCommandCLI) Help() string {
	return `
Usage: seeder init

    Initialize the working directory 'workspace'. Initializes empty deployment plan 'deployment_plan.json' 
and empty state 'deployment_state.json'. Keep in mind that the state is always remote, but the plan is always local.
Copies your 'config.yaml'/'config.yml' file in working directory.
Copies your deployment files in working directory. There are 2 paths for your deployment files:
-   located in a directory called 'deployments', relative to this CLI. The files are valid docker-compose files with extension 'yaml', 'yml'
-   located in the same directory, relative to this CLI. Their name has the form 'deployment_NAME.yaml', 'deployment_NAME.yml'. 

E.g. deployment_mysql8.yml, deployment_cloud_env.yaml

    When you modify a deployment or the CLI configuration, do it outside working directory. If you create a change inside 
working dir, this won't have any effect.

Use 'init' every time you change your configuration or deployments. Use it often.
`
}
