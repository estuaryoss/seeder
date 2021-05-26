package models

type XMetadata struct {
	Replicas int                    `yaml:"replicas,omitempty" validate:"required,min=0"`
	Name     string                 `yaml:"name,omitempty" validate:"required,min=4"`
	Labels   map[string]interface{} `yaml:"labels,omitempty" validate:"required,min=1"`
}

func NewXMetadata() *XMetadata {
	xmetadata := &XMetadata{}
	return xmetadata
}

func (metadata *XMetadata) GetReplicas() int {
	return metadata.Replicas
}

func (metadata *XMetadata) SetReplicas(replicas int) {
	metadata.Replicas = replicas
}

func (metadata *XMetadata) GetName() string {
	return metadata.Name
}

func (metadata *XMetadata) SetName(name string) {
	metadata.Name = name
}

func (metadata *XMetadata) GetLabels() map[string]interface{} {
	return metadata.Labels
}

func (metadata *XMetadata) SetLabels(labels map[string]interface{}) {
	metadata.Labels = labels
}
