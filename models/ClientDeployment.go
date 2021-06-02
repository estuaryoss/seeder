package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"seeder/constants"
	"seeder/utils"
)

type ClientDeployment struct {
	Version   string                 `yaml:"version" validate:"required,min=3"`
	XMetadata *XMetadata             `yaml:"x-metadata,omitempty" validate:"required"`
	Services  map[string]interface{} `yaml:"services" validate:"required,min=1"`
}

type ClientDeploymentHandler struct {
	Validate *validator.Validate
}

func (h ClientDeploymentHandler) ValidateClientDeployments() error {
	deploymentsDir := constants.DEPLOYMENTS_DIR_BEFORE_INIT
	supportedExtensions := []string{"yaml", "yml"}

	filePaths := utils.ListFiles(deploymentsDir, supportedExtensions, true)

	for _, path := range filePaths {
		clientDeployment := &ClientDeployment{}
		fileContent := utils.ReadFile(path)
		err := yaml.Unmarshal(fileContent, &clientDeployment)
		if err = h.Validate.Struct(clientDeployment); err != nil {
			return err
		}
	}

	fmt.Println("Validated deployments: " + fmt.Sprint(filePaths))
	return nil
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

func (deployment *ClientDeployment) GetMetadata() *XMetadata {
	return deployment.XMetadata
}

func (deployment *ClientDeployment) SetMetadata(metadata *XMetadata) {
	deployment.XMetadata = metadata
}

func (deployment *ClientDeployment) GetServices() map[string]interface{} {
	return deployment.Services
}

func (deployment *ClientDeployment) SetServices(services map[string]interface{}) {
	deployment.Services = services
}
