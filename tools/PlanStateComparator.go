package tools

import (
	"reflect"
	"seeder/models"
)

type PlanStateComparator struct {
	PlanDeployments  []*models.ServerDeployment
	StateDeployments []*models.ServerDeployment
	PlannedChanges   []*models.ServerDeployment
	NoChanges        []*models.ServerDeployment
}

func NewPlanStateComparator(plan []*models.ServerDeployment, state []*models.ServerDeployment) *PlanStateComparator {
	return &PlanStateComparator{
		PlanDeployments:  plan,
		StateDeployments: state,
		PlannedChanges:   make([]*models.ServerDeployment, 0),
		NoChanges:        make([]*models.ServerDeployment, 0),
	}
}

func (planStateComparator *PlanStateComparator) GetPlannedChanges() []*models.ServerDeployment {
	for _, planDeployment := range planStateComparator.PlanDeployments {
		if !planStateComparator.isDeploymentFound(planDeployment) {
			planDeployment.OmitDeployment = false
			planStateComparator.PlannedChanges = append(planStateComparator.PlannedChanges, planDeployment)
		}
	}

	return planStateComparator.PlannedChanges
}

func (planStateComparator *PlanStateComparator) GetNoChanges() []*models.ServerDeployment {
	for _, planDeployment := range planStateComparator.PlanDeployments {
		if planStateComparator.isDeploymentFound(planDeployment) {
			planDeployment.OmitDeployment = true
			planStateComparator.NoChanges = append(planStateComparator.NoChanges, planDeployment)
		}
	}

	return planStateComparator.NoChanges
}

func (planStateComparator *PlanStateComparator) GetPlan() []*models.ServerDeployment {
	return planStateComparator.PlanDeployments
}

func (planStateComparator *PlanStateComparator) isDeploymentFound(deployment *models.ServerDeployment) bool {
	for _, stateDeployment := range planStateComparator.StateDeployments {
		if planStateComparator.isDeploymentEqual(deployment, stateDeployment) {
			deployment.HomePageUrl = stateDeployment.HomePageUrl
			return true
		}
	}

	return false
}

func (planStateComparator *PlanStateComparator) isDeploymentEqual(plan *models.ServerDeployment, state *models.ServerDeployment) bool {
	if len(plan.Containers) != len(state.Containers) {
		plan.RecreateDeployment = true
	}
	if plan.Id == state.Id && planStateComparator.isMetadataEqual(plan.Metadata, state.Metadata) &&
		len(plan.Containers) == len(state.Containers) {
		return true
	}

	return false
}

func (planStateComparator *PlanStateComparator) isMetadataEqual(planMetadata *models.XMetadata, stateMetadata *models.XMetadata) bool {
	if planMetadata.File == stateMetadata.File &&
		planMetadata.Name == stateMetadata.Name &&
		reflect.DeepEqual(planMetadata.Labels, stateMetadata.Labels) {

		return true
	}

	return false
}
