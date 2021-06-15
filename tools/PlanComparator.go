package tools

import (
	"reflect"
	"seeder/models"
)

type PlanComparator struct {
	OriginalPlan []*models.ServerDeployment
	NewPlan      []*models.ServerDeployment
	Changes      []*models.ServerDeployment
}

func NewPlanComparator(origPlan []*models.ServerDeployment, newPlan []*models.ServerDeployment) *PlanComparator {
	return &PlanComparator{
		OriginalPlan: origPlan,
		NewPlan:      newPlan,
	}
}

func (planComparator *PlanComparator) GetChanges() []*models.ServerDeployment {
	for _, planDeployment := range planComparator.OriginalPlan {
		if !planComparator.isDeploymentFound(planDeployment) {
			planComparator.Changes = append(planComparator.OriginalPlan, planDeployment)
		}
	}

	return planComparator.Changes
}

func (planComparator *PlanComparator) isDeploymentFound(deployment *models.ServerDeployment) bool {
	for _, newDeployment := range planComparator.NewPlan {
		if planComparator.isDeploymentEqual(deployment, newDeployment) {
			deployment.Deployer = newDeployment.Deployer
			deployment.Discovery = newDeployment.Discovery

			return true
		}
	}

	return false
}

func (planComparator *PlanComparator) isDeploymentEqual(plan *models.ServerDeployment, newPlan *models.ServerDeployment) bool {
	if plan.Id == newPlan.Id && planComparator.isMetadataEqual(plan.Metadata, newPlan.Metadata) &&
		len(plan.Containers) == len(newPlan.Containers) &&
		plan.RecreateDeployment == newPlan.RecreateDeployment {

		return true
	}

	return false
}

func (planComparator *PlanComparator) isMetadataEqual(planMetadata *models.XMetadata, newPlanMetadata *models.XMetadata) bool {
	if planMetadata.File == newPlanMetadata.File &&
		planMetadata.Name == newPlanMetadata.Name &&
		reflect.DeepEqual(planMetadata.Labels, newPlanMetadata.Labels) {

		return true
	}

	return false
}
