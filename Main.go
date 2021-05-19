package main

import (
	"log"
	"os"
	"seeder/factory"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("seeder", "1.0.0")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"init":     factory.Init,
		"validate": factory.Validate,
		"plan":     factory.Plan,
		"apply":    factory.Apply,
		"destroy":  factory.Destroy,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
