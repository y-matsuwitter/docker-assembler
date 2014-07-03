package main

import (
	"encoding/json"
	"github.com/fsouza/go-dockerclient"
	"io/ioutil"
	"log"
	"os"
)

//Config is a Projects configuration.
type Config struct {
	Project    string
	RootImage  string
	Maintainer string
	DockerHost string
}

// OpenConfig opens the configuration file written as json.
func OpenConfig() (config Config) {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		log.Fatalf("File error: %v\n", e)
	}
	// make a read buffer
	e = json.Unmarshal(file, &config)
	if e != nil {
		log.Fatalf("Parse error: %v\n", e)
	}
	return
}

// CreateConfig creates base config.json file.
func CreateConfig() error {
	base := `{
    "Project": "sample",
    "RootImage": "ubuntu:12.04",
    "Maintainer": "none",
    "DockerHost": "unix:///var/run/docker.sock"
}
`
	wr, err := os.Create("config.json")
	if err != nil {
		return err
	}

	defer wr.Close()
	_, err = wr.Write([]byte(base))
	return err
}

func (config Config) GetDockerClient() *docker.Client {
	client, _ := docker.NewClient(config.DockerHost)
	return client
}
