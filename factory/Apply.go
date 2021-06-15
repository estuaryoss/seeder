package factory

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/cli"
	"os"
	"seeder/constants"
	"seeder/models"
	"seeder/services"
	"seeder/tools"
	"seeder/utils"
	"time"
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
	fmt.Println(fmt.Sprintf("Applying the plan found at %s", constants.DEPLOYMENT_PLAN))

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
		pDeployments := getPlannedChanges()
		deploy(yamlConfig, pDeployments)
		savePlan()
		if len(pDeployments) == 0 {
			break
		}
		fmt.Println("Waiting ...")
		time.Sleep(30 * time.Second)
	}

	//save state
	saveRemoteState()

	//save plan
	savePlan()

	return 0
}

func savePlan() {
	plan := getPlan()
	jsonPlan, err := json.Marshal(plan)
	if err != nil {
		fmt.Println(err.Error())
	}
	utils.WriteFile(constants.DEPLOYMENT_PLAN, jsonPlan)
}

func deploy(yamlConfig models.YamlConfig, plannedDeployments []*models.ServerDeployment) {
	for _, deployment := range plannedDeployments {
		if deployment.RecreateDeployment == false {
			continue
		}
		if deployment.Discovery != constants.NA {
			discoveryService := services.NewDiscoveryService(deployment.Discovery, yamlConfig.GetAccessToken())
			discoveryService.DeleteDeploymentId(deployment)
			discoveryService.PostDeploymentUnicast(deployment,
				utils.ReadFile(constants.DEPLOYMENT_DIR_AFTER_INIT+string(os.PathSeparator)+deployment.Metadata.File))
		} else {
			deployerService := services.NewDeployerService(deployment.Deployer, yamlConfig.GetAccessToken())
			deployerService.DeleteDeploymentId(deployment)
			deployerService.PostDeployment(deployment,
				utils.ReadFile(constants.DEPLOYMENT_DIR_AFTER_INIT+string(os.PathSeparator)+deployment.Metadata.File))
		}
	}
}

func getPlannedChanges() []*models.ServerDeployment {
	remoteDeployments := make([]*models.ServerDeployment, 0)
	err := json.Unmarshal(utils.ReadFile(constants.DEPLOYMENT_STATE), &remoteDeployments)
	if err != nil {
		fmt.Println(err.Error())
	}
	deploymentPlanCreator := tools.NewDeploymentPlanCreator(remoteDeployments)

	return deploymentPlanCreator.GetPlannedChanges()
}

func getPlan() []*models.ServerDeployment {
	remoteDeployments := make([]*models.ServerDeployment, 0)
	err := json.Unmarshal(utils.ReadFile(constants.DEPLOYMENT_STATE), &remoteDeployments)
	if err != nil {
		fmt.Println(err.Error())
	}
	deploymentPlanCreator := tools.NewDeploymentPlanCreator(remoteDeployments)

	return deploymentPlanCreator.GetPlan()
}

func saveRemoteState() {
	remoteStateFetcher := tools.NewRemoteStateFetcher()
	remoteStateDeployments := remoteStateFetcher.GetDeployments()

	jsonRemoteStateDeployments, err := json.Marshal(remoteStateDeployments)
	if err != nil {
		fmt.Println(err.Error())
	}
	utils.WriteFile(constants.DEPLOYMENT_STATE, jsonRemoteStateDeployments)
}

func (c *applyCommandCLI) Synopsis() string { return "Usage: seeder apply" }
func (c *applyCommandCLI) Help() string {
	return `
Usage: seeder apply

    Applies the deployment plan. It will bring up all your deployments and sync the current plan with the remote state.
    Saves the remote state as 'deployment_state.json' in 'workspace directory'.

Quote: 'You don't go to war without a plan'

Call it after 'plan', before configuration is applied on remote. Always.

After all deployments are synced:
- it will save a new plan where no deployments are marked for deployment.
- it will save a new remote state with all remote deployments.
`
}
