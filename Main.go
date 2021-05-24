package main

import (
	"log"
	"os"
	"seeder/constants"
	"seeder/factory"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI(constants.APP_NAME, constants.APP_VERSION)
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"validate": factory.Validate,
		"init":     factory.Init,
		"plan":     factory.Plan,
		"apply":    factory.Apply,
		"show":     factory.Show,
		"destroy":  factory.Destroy,
		"version":  factory.Version,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
