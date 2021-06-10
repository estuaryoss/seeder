package tools

import (
	"seeder/constants"
	"seeder/models"
	"seeder/services"
)

type RemoteStateFetcher struct {
	Deployments []interface{}
}

func NewRemoteStateFetcher() *RemoteStateFetcher {
	return &RemoteStateFetcher{Deployments: make([]interface{}, 0)}
}

func (remoteStateFetcher *RemoteStateFetcher) GetDeployments() []interface{} {
	yamlConfig := models.NewYamlConfig().GetYamlConfig()
	infrastructureBuilder := NewInfrastructureBuilder()
	discoveriesMapDeployers := infrastructureBuilder.GetInfrastructureNodes()
	for discoveryHomePageUrl, deployers := range discoveriesMapDeployers {
		if discoveryHomePageUrl != constants.NA {
			//access the deployers through discovery(ies)
			discoveryService := services.NewDiscoveryService(discoveryHomePageUrl, yamlConfig.GetAccessToken())
			for _, deployerHomePageUrl := range deployers {
				discoveryDeployments := discoveryService.GetDeploymentUnicast(deployerHomePageUrl).GetDescription().([]interface{})
				remoteStateFetcher.appendDiscoveryDeployments(discoveryHomePageUrl, deployerHomePageUrl, discoveryDeployments)
			}
		} else {
			//access directly the deployers
			for _, deployerHomePageUrl := range deployers {
				deployerService := services.NewDeployerService(deployerHomePageUrl, yamlConfig.GetAccessToken())
				deployerDeployments := deployerService.HttpClientGetDeployments().GetDescription().([]interface{})
				remoteStateFetcher.appendDeployments(discoveryHomePageUrl, deployerHomePageUrl, deployerDeployments)
			}
		}
	}
	return remoteStateFetcher.Deployments
}

func (remoteStateFetcher *RemoteStateFetcher) appendDeployments(discoveryHomePageUrl string, deployerHomePageUrl string, deployments []interface{}) {
	for _, deployment := range deployments {
		deploymentEnriched := deployment.(map[string]interface{})
		deploymentEnriched["deployer"] = deployerHomePageUrl
		deploymentEnriched["discovery"] = discoveryHomePageUrl
		remoteStateFetcher.Deployments = append(remoteStateFetcher.Deployments, deploymentEnriched)
	}
}

func (remoteStateFetcher *RemoteStateFetcher) appendDiscoveryDeployments(discoveryHomePageUrl string, deployerHomePageUrl string,
	discoveryDeployments []interface{}) {
	for _, discoveryDeployment := range discoveryDeployments {
		deployerDeployments := discoveryDeployment.(map[string]interface{})["description"].([]interface{})
		remoteStateFetcher.appendDeployments(discoveryHomePageUrl, deployerHomePageUrl, deployerDeployments)
	}
}
