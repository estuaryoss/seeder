package tools

import (
	"seeder/models"
)

type DeploymentPlanCreator struct {
	PlanDeployments   []*models.ServerDeployment
	RemoteDeployments []*models.ServerDeployment
}

func NewDeploymentPlanCreator(remoteDeployments []*models.ServerDeployment) *DeploymentPlanCreator {
	return &DeploymentPlanCreator{
		PlanDeployments:   models.NewDeploymentPlan().GetDeploymentPlan(),
		RemoteDeployments: remoteDeployments,
	}
}

func (deploymentPlanCreator *DeploymentPlanCreator) GetPlannedChanges() []*models.ServerDeployment {
	infrastructureBuilder := NewInfrastructureBuilder()
	discoveriesMapDeployers := infrastructureBuilder.GetInfrastructureNodes()

	deploymentPlanPolicy := NewDeploymentPlanPolicy()
	planStateComparator := NewPlanStateComparator(deploymentPlanCreator.PlanDeployments, deploymentPlanCreator.RemoteDeployments)
	plannedChanges := planStateComparator.GetPlannedChanges()

	deploymentPlanPolicy.ApplyPolicy(plannedChanges, discoveriesMapDeployers)

	return plannedChanges
}

func (deploymentPlanCreator *DeploymentPlanCreator) GetNoChanges() []*models.ServerDeployment {
	planStateComparator := NewPlanStateComparator(deploymentPlanCreator.PlanDeployments, deploymentPlanCreator.RemoteDeployments)
	return planStateComparator.GetNoChanges()
}

func (deploymentPlanCreator *DeploymentPlanCreator) GetPlan() []*models.ServerDeployment {
	planStateComparator := NewPlanStateComparator(deploymentPlanCreator.PlanDeployments, deploymentPlanCreator.RemoteDeployments)

	planStateComparator.GetNoChanges()
	planStateComparator.GetPlannedChanges()

	return planStateComparator.GetPlan()
}
