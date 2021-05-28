package factory

import (
	"fmt"
	"github.com/mitchellh/cli"
	"seeder/constants"
	"seeder/utils"
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
	fmt.Println("Initializing workspace ...")
	utils.CreateDir(constants.WORKSPACE)
	utils.CreateDir(constants.DEPLOYMENT_DIR_AFTER_INIT)
	utils.CreateFileIfNotExist(constants.DEPLOYMENT_PLAN)
	utils.CreateFileIfNotExist(constants.DEPLOYMENT_STATE)

	supportedExtensions := []string{"yaml", "yml"}
	filePaths := utils.ListFiles(constants.DEPLOYMENTS_DIR_BEFORE_INIT, supportedExtensions)

	for _, path := range filePaths {
		fileContent := utils.ReadFile(constants.DEPLOYMENTS_DIR_BEFORE_INIT + "/" + path)
		utils.WriteFile(constants.DEPLOYMENT_DIR_AFTER_INIT+"/"+path, fileContent)
	}

	return 0
}

func (c *initCommandCLI) Synopsis() string { return "Usage: seeder init" }
func (c *initCommandCLI) Help() string {
	return `
Usage: seeder init

    Initialize the working directory 'workspace'. Initializes empty deployment plan 'deployment_plan.json' 
and empty state 'deployment_state.json'. Keep in mind that the state is always remote, but the plan is always local.
Copies CLI's global configuration 'config.yaml' file in the working directory.
Copies your deployment files from a directory called 'deployments' in the working directory. The files are valid docker-compose files with extension 'yaml', 'yml'

    When you modify a deployment or the global CLI configuration, do it inside 'deployments' directory. If you create a change inside 
working directory, this change won't have any effect.

Use 'init' every time you change your configuration or deployments. Use it often.
`
}
