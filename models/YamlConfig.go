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
	Eurekas      []string `yaml:"eureka" validate:"required_without_all=Discoveries Deployers,dive,url"`
	Discoveries  []string `yaml:"discovery" validate:"required_without_all=Eurekas Deployers,dive,url"`
	Deployers    []string `yaml:"deployer" validate:"required_without_all=Eurekas Discoveries,dive,url"`
}

type ConfigHandler struct {
	Validate *validator.Validate
}

func NewYamlConfig() *YamlConfig {
	return &YamlConfig{}
}

func (config YamlConfig) GetYamlConfig() YamlConfig {
	yamlConfig := YamlConfig{}

	fileContent := utils.ReadFile(constants.CONFIG_YAML_AFTER_INIT)

	err := yaml.Unmarshal(fileContent, &yamlConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return yamlConfig
}

func (h ConfigHandler) ValidateConfig() error {
	path := constants.CONFIG_YAML

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
	return config.Eurekas
}

func (config *YamlConfig) SetEureka(eureka []string) {
	config.Eurekas = eureka
}

func (config *YamlConfig) GetDiscoveries() []string {
	return config.Discoveries
}

func (config *YamlConfig) SetDiscovery(discovery []string) {
	config.Discoveries = discovery
}

func (config *YamlConfig) GetDeployers() []string {
	return config.Deployers
}

func (config *YamlConfig) SetDeployer(deployer []string) {
	config.Deployers = deployer
}
