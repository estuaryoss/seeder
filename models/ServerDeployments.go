package models

type ServerDeployment struct {
	Id                 string     `json:"id,omitempty"`
	Metadata           *XMetadata `json:"metadata,omitempty"`
	Containers         []string   `json:"containers,omitempty"`
	Discovery          string     `json:"discovery,omitempty"`
	Deployer           string     `json:"deployer,omitempty"`
	RecreateDeployment bool       `json:"recreateDeployment,omitempty"`
}

func NewServerDeployments() *ServerDeployment {
	deployments := &ServerDeployment{}
	return deployments
}

func (deployments *ServerDeployment) GetId() string {
	return deployments.Id
}

func (deployments *ServerDeployment) SetId(id string) {
	deployments.Id = id
}

func (deployments *ServerDeployment) GetMetadata() *XMetadata {
	return deployments.Metadata
}

func (deployments *ServerDeployment) SetMetadata(metadata *XMetadata) {
	deployments.Metadata = metadata
}

func (deployments *ServerDeployment) GetContainers() []string {
	return deployments.Containers
}

func (deployments *ServerDeployment) SetContainers(containers []string) {
	deployments.Containers = containers
}

func (deployments *ServerDeployment) GetDeployer() string {
	return deployments.Deployer
}

func (deployments *ServerDeployment) SetDeployer(deployer string) {
	deployments.Deployer = deployer
}

func (deployments *ServerDeployment) GetDiscovery() string {
	return deployments.Discovery
}

func (deployments *ServerDeployment) SetDiscovery(discovery string) {
	deployments.Discovery = discovery
}
