package factory

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/cli"
	"log"
	"seeder/models"
)

func Validate() (cli.Command, error) {
	validate := &validateCommandCLI{}
	return validate, nil
}

type validateCommandCLI struct {
	Args []string
}

func (c *validateCommandCLI) Run(args []string) int {
	c.Args = args
	fmt.Println("Validating configuration and deployments ...")

	configHandler := models.ConfigHandler{Validate: validator.New()}
	clientDeploymentHandler := models.ClientDeploymentHandler{Validate: validator.New()}

	err := configHandler.ValidateConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = clientDeploymentHandler.ValidateClientDeployments()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return 0
}

func (c *validateCommandCLI) Synopsis() string { return "Usage: seeder validate" }
func (c *validateCommandCLI) Help() string {
	return `
Usage: seeder validate

  Performs several layers of validation before trying to create a deployment plan:
- global config validation
- deployment files validation

  It only validates local files, before being moved in the working directory.
  The 'validate' operation does not perform any validation on remote objects/state.

Call it before 'init'. Always.
`
}
