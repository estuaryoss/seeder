package models

type ClientDeployment struct {
	Version  string                 `yaml:"version"`
	Metadata Metadata               `yaml:"x-metadata"`
	Services map[string]interface{} `yaml:"services"`
}

func NewClientDeployment() *ClientDeployment {
	deployment := &ClientDeployment{}
	return deployment
}

func (deployment *ClientDeployment) GetVersion() string {
	return deployment.Version
}

func (deployment *ClientDeployment) SetVersion(version string) {
	deployment.Version = version
}

func (deployment *ClientDeployment) GetMetadata() Metadata {
	return deployment.Metadata
}

func (deployment *ClientDeployment) SetMetadata(metadata Metadata) {
	deployment.Metadata = metadata
}

func (deployment *ClientDeployment) GetServices() map[string]interface{} {
	return deployment.Services
}

func (deployment *ClientDeployment) SetServices(services map[string]interface{}) {
	deployment.Services = services
}
