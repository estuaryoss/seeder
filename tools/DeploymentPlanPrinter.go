package tools

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"seeder/constants"
	"seeder/models"
	"seeder/utils"
)

type DeploymentPlanPrinter struct {
}

func NewDeploymentPlanPrinter() *DeploymentPlanPrinter {
	deploymentPlanPrinter := &DeploymentPlanPrinter{}
	return deploymentPlanPrinter
}

func (deploymentPlanPrinter *DeploymentPlanPrinter) Print(planFilePath string) error {
	fmt.Println(fmt.Sprintf("Printing plan found at: %s ", constants.DEPLOYMENT_PLAN))

	plan := make([]*models.ServerDeployment, 0)
	err := json.Unmarshal(utils.ReadFile(planFilePath), &plan)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Id", "Deployer(s)", "Container(s)", "Metadata", "Local File"})
	for index, serverDeployment := range plan {
		metadata, _ := json.Marshal(serverDeployment.Metadata)
		t.AppendRow([]interface{}{index, serverDeployment.Id, serverDeployment.HomePageUrl, serverDeployment.Containers,
			string(metadata), serverDeployment.Metadata.File})
	}
	t.Render()

	return nil
}

func (deploymentPlanPrinter *DeploymentPlanPrinter) PrintFromArray(plan []*models.ServerDeployment) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Id", "Deployer(s)", "Container(s)", "Metadata", "Local File"})
	for index, serverDeployment := range plan {
		metadata, _ := json.Marshal(serverDeployment.Metadata)
		t.AppendRow([]interface{}{index, serverDeployment.Id, serverDeployment.HomePageUrl, serverDeployment.Containers,
			string(metadata), serverDeployment.Metadata.File})
	}
	t.Render()
}
