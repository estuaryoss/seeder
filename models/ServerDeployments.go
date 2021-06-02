package models

type ServerDeployment struct {
	Id          string     `json:"id,omitempty"`
	Metadata    *XMetadata `json:"metadata,omitempty"`
	Containers  []string   `json:"containers,omitempty"`
	HomePageUrl string     `json:"homePageUrl,omitempty"`
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

func (deployments *ServerDeployment) GetHomePageUrl() string {
	return deployments.HomePageUrl
}

func (deployments *ServerDeployment) SetHomePageUrl(homePageUrl string) {
	deployments.HomePageUrl = homePageUrl
}
