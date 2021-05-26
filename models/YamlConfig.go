package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"seeder/utils"
)

type YamlConfig struct {
	DeployPolicy string   `yaml:"deploy_policy" validate:"required,oneof=fill robin random"`
	AccessToken  string   `yaml:"access_token" validate:"required,min=4"`
	Eureka       []string `yaml:"eureka" validate:"required_without=discovery deployer,gt=0,dive,url"`
	Discovery    []string `yaml:"discovery" validate:"required_without=eureka deployer,gt=0,dive,url"`
	Deployer     []string `yaml:"deployer" validate:"required_without=eureka discovery,gt=0,dive,url"`
}

type ConfigHandler struct {
	Validate *validator.Validate
}

func (h ConfigHandler) ValidateConfig() error {
	path := "config.yaml"

	yamlConfig := YamlConfig{}

	fileContent := utils.ReadFile(path)

	err := yaml.Unmarshal(fileContent, &yamlConfig)

	if err = h.Validate.Struct(yamlConfig); err != nil {
		return err
	}
	fmt.Println("Validated config: " + path)
	return nil
}

func (config *YamlConfig) GetDeployPolicy() string {
	return config.DeployPolicy
}

func (config *YamlConfig) SetDeployPolicy(deployPolicy string) {
	config.DeployPolicy = deployPolicy
}

func (config *YamlConfig) GetAccessToken() string {
	return config.AccessToken
}

func (config *YamlConfig) SetAccessToken(accessToken string) {
	config.AccessToken = accessToken
}

func (config *YamlConfig) GetEureka() []string {
	return config.Eureka
}

func (config *YamlConfig) SetEureka(eureka []string) {
	config.Eureka = eureka
}

func (config *YamlConfig) GetDiscovery() []string {
	return config.Discovery
}

func (config *YamlConfig) SetDiscovery(discovery []string) {
	config.Discovery = discovery
}

func (config *YamlConfig) GetDeployer() []string {
	return config.Deployer
}

func (config *YamlConfig) SetDeployer(deployer []string) {
	config.Deployer = deployer
}
