package factory

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/cli"
	"os"
	"reflect"
	"seeder/constants"
	"seeder/models"
	"seeder/tools"
	"seeder/utils"
	"strings"
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
	fmt.Println("Saving and printing the deployment newPlan ...")
	deploymentPlanCreator := tools.NewDeploymentPlanCreator()
	newPlan := deploymentPlanCreator.CreateDeploymentPlan()
	deploymentPlanPrinter := tools.NewDeploymentPlanPrinter()
	deploymentPlanPolicy := tools.NewDeploymentPlanPolicy()

	jsonPlan, err := json.Marshal(newPlan)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}
	alreadySavedPlan := make([]*models.ServerDeployment, 0)
	err = json.Unmarshal(utils.ReadFile(constants.DEPLOYMENT_PLAN), &alreadySavedPlan)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}

	if !reflect.DeepEqual(alreadySavedPlan, newPlan) {
		deploymentPlanPrinter.PrintFromArray(newPlan)

		fmt.Println(fmt.Sprintf("Current plan is different from the one found at '%s'. "+
			"Do you want to save the current plan ? [yes/no] : ", constants.DEPLOYMENT_PLAN))
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			return constants.FAILURE
		}
		answer = strings.Replace(answer, "\n", "", -1)
		if strings.Compare(answer, "yes") == 0 {
			utils.WriteFile(constants.DEPLOYMENT_PLAN, jsonPlan)
			fmt.Println(fmt.Sprintf("The current plan was saved at location:  %s.", constants.DEPLOYMENT_PLAN))
		} else {
			fmt.Println("The current plan was discarded.")
		}
	}

	err = deploymentPlanPrinter.Print(constants.DEPLOYMENT_PLAN)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}

	err = deploymentPlanPolicy.PrintWarnings(constants.DEPLOYMENT_PLAN)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}

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
