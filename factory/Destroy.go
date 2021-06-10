package factory

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/cli"
	"seeder/constants"
	"seeder/models"
	"seeder/services"
	"seeder/utils"
	"time"
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
	fmt.Println(fmt.Sprintf("Destroying deployments found in file %s", constants.DEPLOYMENT_PLAN))

	yamlConfig := models.NewYamlConfig().GetYamlConfig()

	plannedDeployments := make([]*models.ServerDeployment, 0)
	err := json.Unmarshal(utils.ReadFile(constants.DEPLOYMENT_PLAN), &plannedDeployments)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}

	//enter check loop
	for plannedDeployments := getPlannedDeployments(); len(plannedDeployments) != 0; {
		fmt.Println("Waiting for all deployments to be destroyed ...")
		destroy(yamlConfig, plannedDeployments)
		saveRemoteState()
		time.Sleep(30 * time.Second)
	}

	//save plan
	plan := getPlan()
	jsonPlan, err := json.Marshal(plan)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}
	utils.WriteFile(constants.DEPLOYMENT_PLAN, jsonPlan)

	return 0
}

func destroy(yamlConfig models.YamlConfig, plannedDeployments []*models.ServerDeployment) {
	for _, deployment := range plannedDeployments {
		if deployment.RecreateDeployment == false {
			continue
		}
		if deployment.Discovery != constants.NA {
			discoveryService := services.NewDiscoveryService(deployment.Discovery, yamlConfig.GetAccessToken())
			discoveryService.DeleteDeploymentId(deployment)
		} else {
			deployerService := services.NewDeployerService(deployment.Deployer, yamlConfig.GetAccessToken())
			deployerService.DeleteDeploymentId(deployment)
		}
	}
}

func (c *destroyCommandCLI) Synopsis() string { return "Usage: seeder destroy" }
func (c *destroyCommandCLI) Help() string {
	return `
Usage: seeder destroy [options] id

    Destroys the remote state and empties your local plan.

Scenarios: 
    Call it without arguments and it will erase all your remote state.
    Call it with args and destroy only the objects specified by 'id'.

At the end of this action:
- new plan will be saved with all the deployments marked for deployment
`
}
