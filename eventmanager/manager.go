package eventmanager

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

const (
	endpoint = "unix:///var/run/docker.sock"
)

// System Monitor represents each active container in system
type SysMonitor struct {
	table     []docker.APIContainers
	connector *docker.Client
}

func (monitor *SysMonitor) refresh() {
	// TODO: Have to update the system monitor
}

// Monitor initializer creates the docker client and the empty monitor struct
func Init() *SysMonitor {
	apiClient, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	containers := make([]docker.APIContainers, 0)
	return &SysMonitor{containers, apiClient}
}

// returns the docker socket connector
func (monitor *SysMonitor) getConnector() *docker.Client {
	return monitor.connector
}

// Listens containers events and updates the system monitor
func Update(monitor *SysMonitor)  {
	dockerClient := monitor.getConnector()

	containers, err := dockerClient.ListContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		fmt.Println("ID:       ", c.ID)
		fmt.Println("Image:    ", c.Image)
		fmt.Println("State:    ", c.State)
		fmt.Println("Status:   ", c.Status)
		fmt.Println("Ports:    ", c.Ports)
		fmt.Println("Networks: ", c.Networks)
	}
}
