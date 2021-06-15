package factory

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/thoas/go-funk"
	"os"
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

	remoteStateFetcher := tools.NewRemoteStateFetcher()
	remoteStateDeployments := remoteStateFetcher.GetDeployments()

	jsonRemoteStateDeployments, err := json.Marshal(remoteStateDeployments)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}
	utils.WriteFile(constants.DEPLOYMENT_STATE, jsonRemoteStateDeployments)
	fmt.Println(fmt.Sprintf("The remote state was saved at location:  %s.", constants.DEPLOYMENT_STATE))

	remoteDeployments := make([]*models.ServerDeployment, 0)
	err = json.Unmarshal(utils.ReadFile(constants.DEPLOYMENT_STATE), &remoteDeployments)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}

	deploymentPlanCreator := tools.NewDeploymentPlanCreator(remoteDeployments)
	plannedChanges := deploymentPlanCreator.GetPlannedChanges()
	noChanges := deploymentPlanCreator.GetNoChanges()
	plan := deploymentPlanCreator.GetPlan()

	deploymentPlanPrinter := tools.NewDeploymentPlanPrinter()
	deploymentPlanPolicy := tools.NewDeploymentPlanPolicy()

	jsonPlan, err := json.Marshal(plan)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}

	if funk.Contains(c.Args, "-auto-approve") {
		utils.WriteFile(constants.DEPLOYMENT_PLAN, jsonPlan)
		fmt.Println(fmt.Sprintf("The current plan was saved at location:  %s.", constants.DEPLOYMENT_PLAN))
	}

	alreadySavedPlan := make([]*models.ServerDeployment, 0)
	err = json.Unmarshal(utils.ReadFile(constants.DEPLOYMENT_PLAN), &alreadySavedPlan)
	if err != nil {
		fmt.Println(err.Error())
		return constants.FAILURE
	}

	fmt.Println("Deployments already created: ")
	deploymentPlanPrinter.PrintFromArray(noChanges)

	fmt.Println("Deployments to be (re)/created: ")
	deploymentPlanPrinter.PrintFromArray(plannedChanges)
	planComparator := tools.NewPlanComparator(alreadySavedPlan, plan)

	if len(planComparator.GetChanges()) != 0 {
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
Usage: seeder plan [Options]

    Creates the deployments plan. By default, creating a plan consists of:
        -  saves the plan locally as 'deployment_plan.json' in the 'workspace' folder
        -  reading remote state and comparing to the current local plan
        -  proposing a set of actions in order to sync remote state with the current plan


Options:

  -auto-approve          Skip interactive approval of plan before applying.

Call it after 'init'. Always.

It can be used also to detect leeway between local plan and remote state. Can be called with 'show' action as well.
`
}
