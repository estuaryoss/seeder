package tools

import (
	"fmt"
	"seeder/constants"
	"seeder/models"
	"seeder/services"
	"sort"
)

type InfrastructureBuilder struct {
}

func NewInfrastructureBuilder() *InfrastructureBuilder {
	return &InfrastructureBuilder{}
}

func (infraBuilder *InfrastructureBuilder) GetInfrastructureNodes() map[string][]string {
	yamlConfig := models.NewYamlConfig().GetYamlConfig()

	eurekas := yamlConfig.GetEurekas()
	discoveries := yamlConfig.GetDiscoveries()
	var deployers []string

	discoveriesMapDeployers := make(map[string][]string)

	if eurekas != nil {
		for _, eureka := range eurekas {
			eureka_service := services.NewEurekaClient(eureka)
			eureka_discoveries := eureka_service.GetDiscoveries()

			for _, discovery := range eureka_discoveries {
				discovery_service := services.NewDiscoveryService(discovery, yamlConfig.GetAccessToken())
				discovery_deployers := discovery_service.HttpClientGetDeployers()
				if discovery_deployers.Description == nil {
					fmt.Sprintf("Unable to get deployer apps from Discoveries: %s",
						fmt.Sprint(discovery))
					continue
				}
				for _, deployer := range discovery_deployers.GetDescription().([]interface{}) {
					deployerCast := deployer.(map[string]interface{})
					deployers = append(deployers, deployerCast["homePageUrl"].(string))
				}

				sort.Strings(deployers)
				discoveriesMapDeployers[discovery] = deployers
			}
		}
	}

	if discoveries != nil {
		for _, discovery := range discoveries {
			discovery_service := services.NewDiscoveryService(discovery, yamlConfig.GetAccessToken())
			discovery_deployers := discovery_service.HttpClientGetDeployers()
			if discovery_deployers.Description == nil {
				fmt.Sprintf("Unable to get deployer apps from Discoveries: %s",
					fmt.Sprint(discovery))
				continue
			}
			for _, deployer := range discovery_deployers.GetDescription().([]interface{}) {
				deployerCast := deployer.(map[string]interface{})
				deployers = append(deployers, deployerCast["homePageUrl"].(string))
			}
			sort.Strings(deployers)
			discoveriesMapDeployers[discovery] = deployers
		}
	}

	if yamlConfig.GetDeployers() != nil {
		sort.Strings(yamlConfig.Deployers)
		discoveriesMapDeployers[constants.NA] = yamlConfig.GetDeployers()
	}

	return discoveriesMapDeployers
}
