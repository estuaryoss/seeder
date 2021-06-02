package services

import (
	"fmt"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"strings"
)

type Eureka struct {
	client *eureka.Client
}

func NewEurekaClient(eurekaServer string) *Eureka {
	return &Eureka{client: GetEurekaClient(eurekaServer)}
}

func (e *Eureka) getApps(appName string) []eureka.Application {
	var eurekaApplications []eureka.Application

	if e.client == nil {
		return nil
	}

	apps, err := e.client.GetApplications()
	if err != nil {
		fmt.Sprintf("Unable to get apps from EurekaServer: %s",
			fmt.Sprint(e.client.GetCluster()))
	}

	for _, app := range apps.Applications {
		if strings.Contains(strings.ToLower(app.Name), appName) {
			eurekaApplications = append(eurekaApplications, app)
		}
	}

	return eurekaApplications
}

func (e *Eureka) GetDiscoveries() []string {
	var discoveries []string
	var eurekaApplications = e.getApps("discovery")

	for _, app := range eurekaApplications {
		for _, instanceInfo := range app.Instances {
			discoveries = append(discoveries, instanceInfo.HomePageUrl)
		}
	}

	return discoveries
}

func GetEurekaClient(eurekaServer string) *eureka.Client {

	if eurekaServer != "" {
		client := eureka.NewClient([]string{eurekaServer})

		return client
	}

	return nil
}
