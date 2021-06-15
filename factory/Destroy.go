package factory

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/cli"
	"seeder/constants"
	"seeder/models"
	"seeder/services"
	"seeder/tools"
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
	for {
		saveRemoteState()
		noChanges := getNoChanges()
		destroy(yamlConfig, noChanges)
		savePlan()
		if len(noChanges) == len(plannedDeployments) {
			break
		}
		fmt.Println("Waiting ...")
		time.Sleep(10 * time.Second)
	}

	//save state
	saveRemoteState()

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
		if deployment.Discovery != constants.NA {
			discoveryService := services.NewDiscoveryService(deployment.Discovery, yamlConfig.GetAccessToken())
			discoveryService.DeleteDeploymentId(deployment)
		} else {
			deployerService := services.NewDeployerService(deployment.Deployer, yamlConfig.GetAccessToken())
			deployerService.DeleteDeploymentId(deployment)
		}
	}
}

func getNoChanges() []*models.ServerDeployment {
	remoteDeployments := make([]*models.ServerDeployment, 0)
	err := json.Unmarshal(utils.ReadFile(constants.DEPLOYMENT_STATE), &remoteDeployments)
	if err != nil {
		fmt.Println(err.Error())
	}
	deploymentPlanCreator := tools.NewDeploymentPlanCreator(remoteDeployments)

	return deploymentPlanCreator.GetNoChanges()
}

func (c *destroyCommandCLI) Synopsis() string { return "Usage: seeder destroy" }
func (c *destroyCommandCLI) Help() string {
	return `
Usage: seeder destroy

    Destroys the remote state and empties your local plan.

Scenarios: 
    Call it without arguments and it will erase all your remote state.
    Call it with args and destroy only the objects specified by 'id'.

After all deployments are destroyed:
- it will save a new plan where the deployments are marked for deployment.
- it will save a new remote state with remaining remote deployments.
`
}
