package factory

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/cli"
	"seeder/constants"
	"seeder/models"
	"seeder/services"
	"seeder/tools"
	"seeder/utils"
	"sort"
)

func Plan() (cli.Command, error) {
	plan := &planCommandCLI{}
	return plan, nil
}

type planCommandCLI struct {
	Args []string
}

func (c *planCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println("Saving and printing the deployment plan ...")
	configHandler := models.ConfigHandler{Validate: validator.New()}

	err := configHandler.ValidateConfig()
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}
	yamlConfig := models.NewYamlConfig().GetYamlConfig()

	eurekas := yamlConfig.GetEurekas()
	discoveries := yamlConfig.GetDiscoveries()
	var deployers []string

	if eurekas != nil {
		for _, eureka := range eurekas {
			eureka_service := services.NewEurekaClient(eureka)
			eureka_deployers := eureka_service.GetDeployers()
			for _, deployer := range eureka_deployers {
				deployers = append(deployers, deployer)
			}
		}
	}

	if discoveries != nil {
		for _, discovery := range discoveries {
			discovery_service := services.NewDiscoveryService(discovery, yamlConfig.GetAccessToken())
			discovery_deployers := discovery_service.HttpClientGetDeployers()
			if discovery_deployers.Description == nil {
				fmt.Sprintf("Unable to get deployer apps from Discovery: %s",
					fmt.Sprint(discovery))
				continue
			}
			for _, deployer := range discovery_deployers.GetDescription().([]interface{}) {
				deployerCast := deployer.(map[string]interface{})
				deployers = append(deployers, deployerCast["homePageUrl"].(string))
			}
		}
	}

	if yamlConfig.GetDeployers() != nil {
		deployers = yamlConfig.GetDeployers()
	}
	sort.Strings(deployers)

	deploymentPlan := models.NewDeploymentPlan()
	deploymentPlanPrinter := tools.NewDeploymentPlanPrinter()
	deploymentPlanPolicy := tools.NewDeploymentPlanPolicy()
	plan := deploymentPlan.CreateDeploymentPlan()
	deploymentPlanPolicy.ApplyPolicy(plan, deployers, yamlConfig)

	jsonPlan, err := json.Marshal(plan)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}
	utils.WriteFile(constants.DEPLOYMENT_PLAN, jsonPlan)

	deploymentPlanPrinter.Print(plan)

	deploymentPlanPolicy.PrintWarnings(plan)

	return constants.SUCCESS
}

func (c *planCommandCLI) Synopsis() string { return "Usage: seeder plan" }
func (c *planCommandCLI) Help() string {
	return `
Usage: seeder plan

    Creates the deployments plan. By default, creating a plan consists of:
        -  saves the plan locally as 'deployment_plan.json' in the 'workspace' folder
        -  reading remote state and comparing to the current local plan
        -  proposing a set of actions in order to sync remote state with the current plan

Call it after 'init'. Always.
`
}
