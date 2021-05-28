package tools

import (
	"encoding/json"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"seeder/models"
)

type DeploymentPlanPrinter struct {
}

func NewDeploymentPlanPrinter() *DeploymentPlanPrinter {
	deploymentPlanPrinter := &DeploymentPlanPrinter{}
	return deploymentPlanPrinter
}

func (deploymentPlanPrinter *DeploymentPlanPrinter) Print(serverDeployments []*models.ServerDeployment) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Id", "Deployer", "Containers", "Metadata", "Local File"})
	for index, serverDeployment := range serverDeployments {
		metadata, _ := json.Marshal(serverDeployment.Metadata)
		t.AppendRow([]interface{}{index, serverDeployment.Id, serverDeployment.HomePageUrl, serverDeployment.Containers,
			string(metadata), serverDeployment.Metadata.File})
	}
	t.Render()

}
