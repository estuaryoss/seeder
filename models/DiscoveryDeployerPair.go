package models

type DiscoveryDeployerPair struct {
	DiscoveryDeployerPair map[string]map[string]int
}

func NewDiscoveryDeployerPair() *DiscoveryDeployerPair {
	return &DiscoveryDeployerPair{
		DiscoveryDeployerPair: make(map[string]map[string]int, 0),
	}
}

func (discoveryDeployerPair *DiscoveryDeployerPair) GetDiscoveryDeployerPair() map[string]map[string]int {
	return discoveryDeployerPair.DiscoveryDeployerPair
}
