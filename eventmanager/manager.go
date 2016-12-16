package eventmanager

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

const (
	Endpoint = "unix:///var/run/docker.sock"
)

func Listen()  {
	client, err := docker.NewClient(Endpoint)
	if err != nil {
		panic(err)
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		fmt.Println("ID:          ", c.ID)
		fmt.Println("Image:       ", c.Image)
		fmt.Println("Created:     ", c.Created)
		fmt.Println("Ports:       ", c.Ports)
		fmt.Println("Networks:    ", c.Networks)
		fmt.Println("Command:     ", c.Command)
	}
}
