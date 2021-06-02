package tools

import (
	"encoding/json"
	"fmt"
	"github.com/thoas/go-funk"
	"log"
	"seeder/constants"
	"seeder/models"
	"seeder/services"
	"seeder/utils"
	"sort"
)

type DeploymentPlanPolicy struct {
	YamlConfig models.YamlConfig
}

func NewDeploymentPlanPolicy() *DeploymentPlanPolicy {
	return &DeploymentPlanPolicy{YamlConfig: models.NewYamlConfig().GetYamlConfig()}
}

func (deploymentPlanPrinter *DeploymentPlanPolicy) ApplyPolicy(plan []*models.ServerDeployment,
	discoveriesMapDeployers map[string][]string) {

	for _, deployers := range discoveriesMapDeployers {
		sort.Strings(deployers)
	}

	var deployerServices []*services.DeployerService
	availableSlots := map[string]int{}

	for _, deployers := range discoveriesMapDeployers {
		for _, deployer := range deployers {
			deployerService := services.NewDeployerService(deployer, deploymentPlanPrinter.YamlConfig.AccessToken)
			deployerServices = append(deployerServices, deployerService)
			availableSlots[deployerService.HomePageUrl] = deployerService.HttpClientGetRemainingSlots()
		}
	}

	switch policy := deploymentPlanPrinter.YamlConfig.DeployPolicy; policy {
	case "fill":
		fill(plan, availableSlots)
		fmt.Println(fmt.Sprintf("Applying '%s' deployment policy", constants.FILL))
	case "robin":
		robin(plan, availableSlots)
		fmt.Println(fmt.Sprintf("Applying '%s' deployment policy", constants.ROBIN))
	default:
		robin(plan, availableSlots)
		fmt.Println(fmt.Sprintf("Applying '%s' deployment policy", constants.ROBIN))
	}
}

func fill(plan []*models.ServerDeployment, deployersSlots map[string]int) {
	slotIndex := 0
	for deployerHomeUrl, noSlots := range deployersSlots {
		for i := 0; i < noSlots; i++ {
			plan[slotIndex].HomePageUrl = deployerHomeUrl
			if slotIndex == len(plan)-1 {
				break
			}
			slotIndex++
		}
		if slotIndex == len(plan)-1 {
			break
		}
	}
}

func robin(plan []*models.ServerDeployment, deployersSlots map[string]int) {
	slotIndex := 0
	deployers := utils.GetKeysFromMapInt(deployersSlots)

	for _, noSlots := range deployersSlots {
		for i := 0; i < noSlots; i++ {
			plan[slotIndex].HomePageUrl = deployers[slotIndex%len(deployers)]
			if slotIndex == len(plan)-1 {
				break
			}
			slotIndex++
		}
		if slotIndex == len(plan)-1 {
			break
		}
	}
}

func (deploymentPlanPrinter *DeploymentPlanPolicy) PrintWarnings(planFilePath string) error {
	plan := make([]*models.ServerDeployment, 0)
	err := json.Unmarshal(utils.ReadFile(planFilePath), &plan)
	if err != nil {
		return err
	}

	notDeployableDeployments := funk.Filter(plan, doesNotHaveHomePageUrl)
	l := len(notDeployableDeployments.([]*models.ServerDeployment))
	if l != 0 {
		log.Println(fmt.Sprintf("WARNING: No slots available for %d deployment(s)", l))
	}

	return nil
}

func doesNotHaveHomePageUrl(deployment *models.ServerDeployment) bool {
	return deployment.HomePageUrl == ""
}
