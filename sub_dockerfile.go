package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// SubDockerfile is a data for sub directory
type SubDockerfile struct {
	Name   string
	Config Config
}

func (sd SubDockerfile) createSubDirectory() error {
	return os.Mkdir(sd.Name, 0755)
}

func (sd SubDockerfile) putDockerfile() error {
	tmpl, err := template.New(sd.Name).Parse(`FROM {{.RootImage}}
MAINTAINER {{.Maintainer}}`)
	if err != nil {
		return err
	}
	wr, err := os.Create(sd.Name + "/Dockerfile")
	if err != nil {
		return err
	}

	defer wr.Close()
	return tmpl.Execute(wr, sd.Config)
}

//Create creates a dockerfile and sub-directory.
func (sd SubDockerfile) Create() error {
	err := sd.createSubDirectory()
	if err != nil {
		return err
	}
	return sd.putDockerfile()
}

func (sd SubDockerfile) applyFrom(from string) error {
	rawFile, err := os.Open("Dockerfile.raw")
	if err != nil {
		return err
	}
	defer rawFile.Close()
	reader := bufio.NewReader(rawFile)
	fo, err := os.Create("Dockerfile")
	if err != nil {
		return err
	}
	defer fo.Close()
	writer := bufio.NewWriter(fo)

	line, isPrefix, err := reader.ReadLine()
	for err == nil && !isPrefix {
		s := string(line)
		if strings.HasPrefix(s, "FROM ") {
			fromLine := fmt.Sprintf("FROM %s\n", from)
			writer.Write([]byte(fromLine))
		} else {
			writer.WriteString(s + "\n")
		}
		line, isPrefix, err = reader.ReadLine()
	}
	if err = writer.Flush(); err != nil {
		return err
	}
	return nil
}

// Build builds docker image with "from" image.
func (sd SubDockerfile) Build(from string) (imageName string, err error) {
	imageName = strings.Join([]string{sd.Config.Project, sd.Name}, "-")
	os.Chdir(sd.Name)
	os.Rename("Dockerfile", "Dockerfile.raw")
	sd.applyFrom(from)
	log.Println("Building: ", sd.Name)
	cmd := exec.Command("docker", "-H", sd.Config.DockerHost, "build", "-t", imageName, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	os.Remove("Dockerfile")
	os.Rename("Dockerfile.raw", "Dockerfile")
	os.Chdir("..")
	if err != nil {
		log.Fatal(err)
	}
	return
}
