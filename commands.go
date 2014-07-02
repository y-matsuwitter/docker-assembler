package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

// Commands is a command set of docker-assembler.
var Commands = []cli.Command{
	commandCreate,
	commandAdd,
	commandBuild,
}

var commandCreate = cli.Command{
	Name:  "create",
	Usage: "Setup docker assembler project.",
	Flags: []cli.Flag{
		cli.StringFlag{"s", "sub1,sub2,...", "Sub Docker directory."},
	},
	Description: `
`,
	Action: doCreate,
}

var commandAdd = cli.Command{
	Name:  "add",
	Usage: "Add new sub-Dockerfile and its directory.",
	Description: `
`,
	Action: doAdd,
}

var commandBuild = cli.Command{
	Name:  "build",
	Usage: "Build a Docker image with specified sub-Dockerfiles.",
	Description: `
`,
	Action: doBuild,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doCreate(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("Usage: docker-assembler create [Project name]")
		return
	}
	name := c.Args()[0]
	fmt.Println("Created: " + name)
	for _, sub := range c.Args()[1:] {
		addSubDockerfile(sub, name)
	}
}

func doAdd(c *cli.Context) {
}

func addSubDockerfile(name, basename string) {
	fmt.Println("Created: " + basename + "/" + name)
}

func doBuild(c *cli.Context) {
}
