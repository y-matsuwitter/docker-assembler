package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-assembler"
	app.Version = Version
	app.Usage = "Docker image builder with sub-Dockerfiles"
	app.Author = "y-matsuwitter"
	app.Email = "y.matsu.on.twitter@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
