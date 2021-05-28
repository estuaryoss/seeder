package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"log"
	"seeder/constants"
	"seeder/utils"
)

type YamlConfig struct {
	DeployPolicy string   `yaml:"deploy_policy" validate:"required,oneof=fill robin random"`
	AccessToken  string   `yaml:"access_token" validate:"required,min=4"`
	Eureka       []string `yaml:"eureka" validate:"required_without_all=Discovery Deployer,dive,url"`
	Discovery    []string `yaml:"discovery" validate:"required_without_all=Eureka Deployer,dive,url"`
	Deployer     []string `yaml:"deployer" validate:"required_without_all=Eureka Discovery,dive,url"`
}

type ConfigHandler struct {
	Validate *validator.Validate
}

func NewYamlConfig() *YamlConfig {
	return &YamlConfig{}
}

func (config YamlConfig) GetYamlConfig() YamlConfig {
	yamlConfig := YamlConfig{}

	fileContent := utils.ReadFile(constants.CONFIG_YAML)

	err := yaml.Unmarshal(fileContent, &yamlConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return yamlConfig
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

func (config *YamlConfig) GetEurekas() []string {
	return config.Eureka
}

func (config *YamlConfig) SetEureka(eureka []string) {
	config.Eureka = eureka
}

func (config *YamlConfig) GetDiscoveries() []string {
	return config.Discovery
}

func (config *YamlConfig) SetDiscovery(discovery []string) {
	config.Discovery = discovery
}

func (config *YamlConfig) GetDeployers() []string {
	return config.Deployer
}

func (config *YamlConfig) SetDeployer(deployer []string) {
	config.Deployer = deployer
}
