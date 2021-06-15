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

	yamlConfig := models.NewYamlConfig().GetYamlConfig()
	availableSlots := models.NewDiscoveryDeployerPair().GetDiscoveryDeployerPair()

	for discoveryHomePageUrl, deployers := range discoveriesMapDeployers {
		if discoveryHomePageUrl != constants.NA {
			//access the deployers through discovery(ies)
			discoveryService := services.NewDiscoveryService(discoveryHomePageUrl, yamlConfig.GetAccessToken())
			for _, deployerHomePageUrl := range deployers {
				deployerSlots := make(map[string]int, 0)
				deployerSlots[deployerHomePageUrl] = discoveryService.GetRemainingSlots(deployerHomePageUrl)
				availableSlots[discoveryHomePageUrl] = deployerSlots
			}
		} else {
			//access directly the deployers
			for _, deployerHomePageUrl := range deployers {
				deployerService := services.NewDeployerService(deployerHomePageUrl, deploymentPlanPrinter.YamlConfig.AccessToken)
				deployerSlots := make(map[string]int, 0)
				deployerSlots[deployerHomePageUrl] = deployerService.HttpClientGetRemainingSlots()
				availableSlots[constants.NA] = deployerSlots
			}
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

func fill(plan []*models.ServerDeployment, discoveryDeployerSlots map[string]map[string]int) {
	slotIndex := 0
	if len(plan) == 0 { //in case already deployed
		return
	}
	for discoveryHomePageUrl, deployerSlots := range discoveryDeployerSlots {
		for deployerHomeUrl, noSlots := range deployerSlots {
			for i := 0; i < noSlots; i++ {
				plan[slotIndex].Discovery = discoveryHomePageUrl
				plan[slotIndex].Deployer = deployerHomeUrl
				if slotIndex == len(plan)-1 {
					break
				}
				slotIndex++
			}
		}
		if slotIndex == len(plan)-1 {
			break
		}
	}
}

func robin(plan []*models.ServerDeployment, discoveryDeployerSlots map[string]map[string]int) {
	slotIndex := 0
	for discoveryHomePageUrl, deployerSlots := range discoveryDeployerSlots {
		deployers := utils.GetKeysFromMapInt(deployerSlots)
		for _, noSlots := range deployerSlots {
			for i := 0; i < noSlots; i++ {
				plan[slotIndex].Discovery = discoveryHomePageUrl
				plan[slotIndex].Deployer = deployers[slotIndex%len(deployers)]
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
	return deployment.Deployer == ""
}
