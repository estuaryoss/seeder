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
	ServerDeployments []*models.ServerDeployments
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

func (service *DeployerService) GetApiResponse() *models.ApiResponse {
	return service.ApiResponse
}

func (service *DeployerService) GetServerDeployments() []*models.ServerDeployments {
	return service.ServerDeployments
}

func (service *DeployerService) GetHttpClient() *http.Client {
	return service.HttpClient
}

func (service *DeployerService) HttpClientGetDeployments(homePageUrl string) *models.ApiResponse {
	req, err := http.NewRequest("GET", homePageUrl+"/deployments", nil)
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
	json.Unmarshal(bodyBytes, service.ApiResponse)

	return service.ApiResponse
}

func (service *DeployerService) HttpClientGetDeploymentId(homePageUrl string, id string) *models.ApiResponse {
	req, err := http.NewRequest("GET", homePageUrl+"/deployments/"+id, nil)
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
	json.Unmarshal(bodyBytes, service.ApiResponse)

	return service.ApiResponse
}

func (service *DeployerService) HttpClientPostDeployments(homePageUrl string, body []byte) *models.ApiResponse {
	req, err := http.NewRequest("POST", homePageUrl+"/deployments", bytes.NewBuffer(body))
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
	json.Unmarshal(bodyBytes, service.ApiResponse)

	return service.ApiResponse
}

func (service *DeployerService) HttpClientDeleteDeployments(homePageUrl string) *models.ApiResponse {
	req, err := http.NewRequest("DELETE", homePageUrl+"/deployments", nil)
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
	json.Unmarshal(bodyBytes, service.ApiResponse)

	return service.ApiResponse
}

func (service *DeployerService) HttpClientDeleteDeploymentId(homePageUrl string, id string) *models.ApiResponse {
	req, err := http.NewRequest("DELETE", homePageUrl+"/deployments/"+id, nil)
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
	json.Unmarshal(bodyBytes, service.ApiResponse)

	return service.ApiResponse
}
