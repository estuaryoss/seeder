package models

type ServerDeployments struct {
	Id          string    `json:"id,omitempty"`          //from deployer
	Metadata    XMetadata `json:"metadata,omitempty"`    //from deployer
	Containers  []string  `json:"containers,omitempty"`  //from deployer
	IpPort      string    `json:"ip_port,omitempty"`     //from discovery
	HomePageUrl string    `json:"homePageUrl,omitempty"` //from discovery
}

func NewServerDeployments() *ServerDeployments {
	deployments := &ServerDeployments{}
	return deployments
}

func (deployments *ServerDeployments) GetId() string {
	return deployments.Id
}

func (deployments *ServerDeployments) SetId(id string) {
	deployments.Id = id
}

func (deployments *ServerDeployments) GetMetadata() XMetadata {
	return deployments.Metadata
}

func (deployments *ServerDeployments) SetMetadata(metadata XMetadata) {
	deployments.Metadata = metadata
}

func (deployments *ServerDeployments) GetContainers() []string {
	return deployments.Containers
}

func (deployments *ServerDeployments) SetContainers(containers []string) {
	deployments.Containers = containers
}

func (deployments *ServerDeployments) GetIpPort() string {
	return deployments.IpPort
}

func (deployments *ServerDeployments) SetIpPort(ipPort string) {
	deployments.IpPort = ipPort
}

func (deployments *ServerDeployments) GetHomePageUrl() string {
	return deployments.HomePageUrl
}

func (deployments *ServerDeployments) SetHomePageUrl(homePageUrl string) {
	deployments.HomePageUrl = homePageUrl
}
