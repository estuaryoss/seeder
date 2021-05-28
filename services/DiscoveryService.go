package services

import (
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

func (service *DiscoveryService) HttpClientGetDeployers() *models.ApiResponse {
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
