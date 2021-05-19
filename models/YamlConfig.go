package models

type YamlConfig struct {
	DeployPolicy string   `yaml:"deploy_policy"`
	AccessToken  string   `yaml:"access_token"`
	Eureka       []string `yaml:"eureka"`
	Discovery    []string `yaml:"discovery"`
	Deployer     []string `yaml:"deployer"`
}

func NewYamlConfig() *YamlConfig {
	config := &YamlConfig{}
	return config
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
