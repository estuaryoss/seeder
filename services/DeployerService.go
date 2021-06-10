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

type DeployerService struct {
	HomePageUrl       string
	AccessToken       string
	ApiResponse       *models.ApiResponse
	ServerDeployments []*models.ServerDeployment
	HttpClient        *http.Client
}

func NewDeployerService(homePageUrl string, accessToken string) *DeployerService {
	service := &DeployerService{}
	service.HttpClient = &http.Client{Timeout: time.Minute}
	service.HomePageUrl = homePageUrl
	service.AccessToken = accessToken

	return service
}

func (service *DeployerService) GetHomePageUrl() string {
	return service.HomePageUrl
}

func (service *DeployerService) GetAccessToken() string {
	return service.AccessToken
}

func (service *DeployerService) GetApiResponse() *models.ApiResponse {
	return service.ApiResponse
}

func (service *DeployerService) GetServerDeployments() []*models.ServerDeployment {
	return service.ServerDeployments
}

func (service *DeployerService) GetHttpClient() *http.Client {
	return service.HttpClient
}

func (service *DeployerService) HttpClientGetDeployments() *models.ApiResponse {
	req, err := http.NewRequest("GET", service.HomePageUrl+"/deployments", nil)
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

func (service *DeployerService) HttpClientGetDeploymentId(id string) *models.ApiResponse {
	req, err := http.NewRequest("GET", service.HomePageUrl+"/deployments/"+id, nil)
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

func (service *DeployerService) HttpClientGetRemainingSlots() int {
	maxDeployments := service.HttpClientGetEnvInit().GetDescription().(map[string]interface{})["MAX_DEPLOYMENTS"].(float64)

	return int(maxDeployments) - len(service.HttpClientGetDeployments().GetDescription().([]interface{}))
}

func (service *DeployerService) HttpClientGetEnvInit() *models.ApiResponse {
	req, err := http.NewRequest("GET", service.HomePageUrl+"/envinit", nil)
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
	json.Unmarshal(bodyBytes, apiResponse)

	return service.ApiResponse
}

func (service *DeployerService) PostDeployment(deployment *models.ServerDeployment, deploymentFileContent []byte) *models.ApiResponse {
	req, err := http.NewRequest("POST", service.HomePageUrl+"/deployments", bytes.NewBuffer(deploymentFileContent))
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Token", service.AccessToken)
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
	json.Unmarshal(bodyBytes, apiResponse)

	return service.ApiResponse
}

func (service *DeployerService) DeleteDeployments() *models.ApiResponse {
	req, err := http.NewRequest("DELETE", service.HomePageUrl+"/deployments", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
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
	json.Unmarshal(bodyBytes, apiResponse)

	return service.ApiResponse
}

func (service *DeployerService) DeleteDeploymentId(deployment *models.ServerDeployment) *models.ApiResponse {
	req, err := http.NewRequest("DELETE", service.HomePageUrl+"/deployments/"+deployment.Id, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
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
	json.Unmarshal(bodyBytes, apiResponse)

	return service.ApiResponse
}
