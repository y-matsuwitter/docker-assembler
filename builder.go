package main

import (
	"github.com/fsouza/go-dockerclient"
	"log"
)

type Builder struct {
	Config         Config
	SubDockerfiles []*SubDockerfile
}

func (builder *Builder) BuildImage(imageName string) {
	from := builder.Config.RootImage
	var err error
	for _, sub := range builder.SubDockerfiles {
		from, err = sub.Build(from)
		if err != nil {
			log.Fatal(err)
		}
	}
	client := builder.Config.GetDockerClient()
	client.TagImage(from, docker.TagImageOptions{imageName, true})
	log.Println("Built all images. Name:", imageName)
}
