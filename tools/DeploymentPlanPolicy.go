package tools

import (
	"fmt"
	"github.com/thoas/go-funk"
	"log"
	"math/rand"
	"seeder/constants"
	"seeder/models"
	"seeder/services"
	"seeder/utils"
	"sort"
	"time"
)

type DeploymentPlanPolicy struct {
}

func NewDeploymentPlanPolicy() *DeploymentPlanPolicy {
	return &DeploymentPlanPolicy{}
}

func (deploymentPlanPrinter *DeploymentPlanPolicy) ApplyPolicy(plan []*models.ServerDeployment,
	deployers []string, yamlConfig models.YamlConfig) {
	sort.Strings(deployers)
	var deployerServices []*services.DeployerService
	availableSlots := map[string]int{}

	for _, deployer := range deployers {
		deployerService := services.NewDeployerService(deployer, yamlConfig.AccessToken)
		deployerServices = append(deployerServices, deployerService)
		availableSlots[deployerService.HomePageUrl] = deployerService.HttpClientGetRemainingSlots()
	}

	switch policy := yamlConfig.DeployPolicy; policy {
	case "fill":
		fill(plan, availableSlots)
		fmt.Println(fmt.Sprintf("Applying '%s' deployment policy", constants.FILL))
	case "robin":
		robin(plan, availableSlots)
		fmt.Println(fmt.Sprintf("Applying '%s' deployment policy", constants.ROBIN))
	default:
		random(plan, availableSlots)
		fmt.Println(fmt.Sprintf("Applying '%s' deployment policy", constants.RANDOM))
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

//TODO
func random(plan []*models.ServerDeployment, deployersSlots map[string]int) {
	slotIndex := 0
	deployers := utils.GetKeysFromMapInt(deployersSlots)
	rand.Seed(time.Now().UnixNano())

	for _, noSlots := range deployersSlots {
		for i := 0; i < noSlots; i++ {
			r := rand.Intn(len(deployers)-0) + 0
			plan[slotIndex].HomePageUrl = deployers[r]
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

func (deploymentPlanPrinter *DeploymentPlanPolicy) PrintWarnings(plan []*models.ServerDeployment) {
	undeployableDeployments := funk.Filter(plan, doesNotHaveHomePageUrl)
	l := len(undeployableDeployments.([]*models.ServerDeployment))
	if l != 0 {
		log.Println(fmt.Sprintf("WARNING: No slots available for %d deployment(s)", l))
	}
}

func doesNotHaveHomePageUrl(deployment *models.ServerDeployment) bool {
	return deployment.HomePageUrl == ""
}
