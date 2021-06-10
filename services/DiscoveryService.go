package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"seeder/models"
	"time"
)

type DiscoveryService struct {
	HomePageUrl string
	AccessToken string
	ApiResponse *models.ApiResponse
	HttpClient  *http.Client
}

func NewDiscoveryService(homePageUrl string, accessToken string) *DiscoveryService {
	service := &DiscoveryService{}
	service.HttpClient = &http.Client{Timeout: time.Minute}
	service.HomePageUrl = homePageUrl
	service.AccessToken = accessToken

	return service
}

func (service *DiscoveryService) GetHomePageUrl() string {
	return service.HomePageUrl
}

func (service *DiscoveryService) GetAccessToken() string {
	return service.AccessToken
}

func (service *DiscoveryService) GetApiResponse() *models.ApiResponse {
	return service.ApiResponse
}

func (service *DiscoveryService) GetHttpClient() *http.Client {
	return service.HttpClient
}

func (service *DiscoveryService) GetDeployers() *models.ApiResponse {
	req, err := http.NewRequest("GET", service.HomePageUrl+"/eurekaapps/deployer", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", service.AccessToken)
	resp, err := service.HttpClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	apiResponse := &service.ApiResponse
	err = json.Unmarshal(bodyBytes, apiResponse)

	if err != nil {
		fmt.Print(err.Error())
	}

	return service.ApiResponse
}

func (service *DiscoveryService) GetRemainingSlots(deployerHomePageUrl string) int {
	maxDeploymentsDiscoveryResponse := service.GetEnvInit(deployerHomePageUrl).GetDescription().([]interface{})
	var maxDeployments float64
	for _, envInit := range maxDeploymentsDiscoveryResponse {
		maxDeployments = envInit.(map[string]interface{})["description"].(map[string]interface{})["MAX_DEPLOYMENTS"].(float64)
	}
	discoveryDeployerDeployments := service.GetDeploymentUnicast(deployerHomePageUrl).GetDescription().([]interface{})
	deployerDeployments := make([]interface{}, 0)
	for _, discoveryDeployment := range discoveryDeployerDeployments {
		deployerDeployments = discoveryDeployment.(map[string]interface{})["description"].([]interface{})
	}
	return int(maxDeployments) - len(deployerDeployments)
}

func (service *DiscoveryService) GetEnvInit(deployerHomePageUrl string) *models.ApiResponse {
	req, err := http.NewRequest("GET", service.HomePageUrl+"/deployers/envinit", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", service.AccessToken)
	req.Header.Add("HomePageUrl", deployerHomePageUrl)
	resp, err := service.HttpClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	apiResponse := &service.ApiResponse
	err = json.Unmarshal(bodyBytes, apiResponse)
	if err != nil {
		fmt.Print(err.Error())
	}

	return service.ApiResponse
}

func (service *DiscoveryService) GetDeploymentUnicast(deployerHomePageUrl string) *models.ApiResponse {
	req, err := http.NewRequest("GET", service.HomePageUrl+"/deployers/deployments", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", service.AccessToken)
	req.Header.Add("HomePageUrl", deployerHomePageUrl)
	resp, err := service.HttpClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	apiResponse := &service.ApiResponse
	err = json.Unmarshal(bodyBytes, apiResponse)
	if err != nil {
		fmt.Print(err.Error())
	}

	return service.ApiResponse
}

func (service *DiscoveryService) DeleteDeploymentId(deployment *models.ServerDeployment) *models.ApiResponse {
	req, err := http.NewRequest("DELETE", service.HomePageUrl+"/deployers/deployments/"+deployment.Id, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", service.AccessToken)
	req.Header.Add("HomePageUrl", deployment.Deployer)
	resp, err := service.HttpClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	apiResponse := &service.ApiResponse
	err = json.Unmarshal(bodyBytes, apiResponse)
	if err != nil {
		fmt.Print(err.Error())
	}

	return service.ApiResponse
}

func (service *DiscoveryService) PostDeploymentUnicast(deployment *models.ServerDeployment, deploymentFile []byte) *models.ApiResponse {
	req, err := http.NewRequest("POST", service.HomePageUrl+"/deployers/deployments", bytes.NewBuffer(deploymentFile))
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", service.AccessToken)
	req.Header.Add("HomePageUrl", deployment.Deployer)
	req.Header.Add("Deployment-Id", deployment.Id)

	resp, err := service.HttpClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	apiResponse := &service.ApiResponse
	err = json.Unmarshal(bodyBytes, apiResponse)
	if err != nil {
		fmt.Print(err.Error())
	}

	return service.ApiResponse
}
