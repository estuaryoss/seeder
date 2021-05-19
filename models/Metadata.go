package models

type Metadata struct {
	Replicas int                    `yaml:"replicas"`
	Name     string                 `yaml:"name"`
	Labels   map[string]interface{} `yaml:"labels"`
}

func NewMetadata() *Metadata {
	metadata := &Metadata{}
	return metadata
}

func (metadata *Metadata) GetReplicas() int {
	return metadata.Replicas
}

func (metadata *Metadata) SetReplicas(replicas int) {
	metadata.Replicas = replicas
}

func (metadata *Metadata) GetName() string {
	return metadata.Name
}

func (metadata *Metadata) SetName(name string) {
	metadata.Name = name
}

func (metadata *Metadata) GetLabels() map[string]interface{} {
	return metadata.Labels
}

func (metadata *Metadata) SetLabels(labels map[string]interface{}) {
	metadata.Labels = labels
}
