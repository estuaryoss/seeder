package tools

import (
	"seeder/models"
)

type DeploymentPlanCreator struct {
}

func NewDeploymentPlanCreator() *DeploymentPlanCreator {
	return &DeploymentPlanCreator{}
}

func (deploymentPlanCreator *DeploymentPlanCreator) CreateDeploymentPlan() []*models.ServerDeployment {
	infrastructureBuilder := NewInfrastructureBuilder()
	discoveriesMapDeployers := infrastructureBuilder.GetInfrastructureNodes()

	deploymentPlan := models.NewDeploymentPlan()
	deploymentPlanPolicy := NewDeploymentPlanPolicy()
	plan := deploymentPlan.GetDeploymentPlan()
	deploymentPlanPolicy.ApplyPolicy(plan, discoveriesMapDeployers)

	return plan
}
