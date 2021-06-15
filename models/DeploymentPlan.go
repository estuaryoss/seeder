package models

import (
	"gopkg.in/yaml.v3"
	"os"
	"seeder/constants"
	"seeder/utils"
	"strconv"
)

type DeploymentPlan struct {
	Plan  []*ServerDeployment
	Index int
}

func NewDeploymentPlan() *DeploymentPlan {
	return &DeploymentPlan{Index: 1000}
}

func (deploymentPlan *DeploymentPlan) GetPlan() []*ServerDeployment {
	return deploymentPlan.Plan
}

func (deploymentPlan *DeploymentPlan) GetDeploymentPlan() []*ServerDeployment {
	deploymentPlan.Plan = []*ServerDeployment{}
	fileDeploymentList := utils.ListFiles(constants.DEPLOYMENT_DIR_AFTER_INIT, []string{"yaml", "yml"}, false)

	for _, filePath := range fileDeploymentList {
		clientDeployment := &ClientDeployment{}
		yaml.Unmarshal(utils.ReadFile(constants.DEPLOYMENT_DIR_AFTER_INIT+string(os.PathSeparator)+filePath), clientDeployment)
		for i := 0; i < clientDeployment.XMetadata.Replicas; i++ {
			serverDeployment := &ServerDeployment{}
			deploymentId := strconv.Itoa(deploymentPlan.Index)
			serverDeployment.SetId(deploymentId)
			deploymentPlan.Index++
			serverDeployment.SetContainers(utils.GetKeysFromMap(clientDeployment.GetServices()))
			clientDeployment.XMetadata.File = filePath
			serverDeployment.SetMetadata(clientDeployment.GetMetadata())

			deploymentPlan.Plan = append(deploymentPlan.Plan, serverDeployment)
		}
	}

	return deploymentPlan.Plan
}
