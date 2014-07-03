package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
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
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(name)
	err = CreateConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created: " + name)
}

func doAdd(c *cli.Context) {
	for _, sub := range c.Args() {
		addSubDockerfile(sub)
	}
}

func addSubDockerfile(name string) {
	config := OpenConfig()
	sub := SubDockerfile{Name: name, Config: config}
	err := sub.Create()
	if err != nil {
		fmt.Println("Added: " + name)
	}
}

func doBuild(c *cli.Context) {
	config := OpenConfig()
	builder := Builder{Config: config, SubDockerfiles: []*SubDockerfile{}}
	imageName := c.Args()[0]
	log.Println("ImageName: ", imageName)
	log.Println("Build ", strings.Join(c.Args()[1:], " -> "))
	for _, name := range c.Args()[1:] {
		sub := SubDockerfile{Name: name, Config: config}
		builder.SubDockerfiles = append(builder.SubDockerfiles, &sub)
	}
	builder.BuildImage(imageName)
	log.Println("Completed.")
}
