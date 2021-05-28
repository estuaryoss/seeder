package models

import (
	"gopkg.in/yaml.v3"
	"seeder/constants"
	"seeder/utils"
)

type DeploymentPlan struct {
	Plan []*ServerDeployment
}

func NewDeploymentPlan() *DeploymentPlan {
	return &DeploymentPlan{}
}

func (deploymentPlan *DeploymentPlan) GetPlan() []*ServerDeployment {
	return deploymentPlan.Plan
}

func (deploymentPlan *DeploymentPlan) CreateDeploymentPlan() []*ServerDeployment {
	deploymentPlan.Plan = []*ServerDeployment{}
	fileDeploymentList := utils.ListFiles(constants.DEPLOYMENT_DIR_AFTER_INIT, []string{"yaml", "yml"})

	for _, filePath := range fileDeploymentList {
		clientDeployment := &ClientDeployment{}
		yaml.Unmarshal(utils.ReadFile(constants.DEPLOYMENT_DIR_AFTER_INIT+"/"+filePath), clientDeployment)
		for i := 0; i < clientDeployment.XMetadata.Replicas; i++ {
			serverDeployment := &ServerDeployment{}
			serverDeployment.SetId(utils.GenerateRandomId(8))
			serverDeployment.SetContainers(utils.GetKeysFromMap(clientDeployment.GetServices()))
			clientDeployment.XMetadata.File = filePath
			serverDeployment.SetMetadata(clientDeployment.GetMetadata())

			deploymentPlan.Plan = append(deploymentPlan.Plan, serverDeployment)
		}
	}

	return deploymentPlan.Plan
}
