package eventmanager

import (
    "fmt"
    "strings"

    "github.com/fsouza/go-dockerclient"
    "github.com/gninjava/registryd/storage"
)

const (
    endpoint    = "unix:///var/run/docker.sock"
    instanceKey = "nginx"
)

// System Monitor represents each active container in system
type SysMonitor struct {
    table     []docker.APIContainers
    connector *docker.Client
}

// Init serves as monitor initializer that
// creates the docker client and the empty monitor struct
func Init() *SysMonitor {
    apiClient, err := docker.NewClient(endpoint)
    if err != nil {
        panic(err)
    }

    containers := make([]docker.APIContainers, 0)
    return &SysMonitor{containers, apiClient}
}

// getConnector returns the docker socket connector
func (monitor *SysMonitor) getConnector() *docker.Client {
    return monitor.connector
}

// Update listens containers' events and updates the system monitor
func Update(monitor *SysMonitor)  {
    dockerClient := monitor.getConnector()

    containers, err := dockerClient.ListContainers(docker.ListContainersOptions{All: true})
    if err != nil {
        panic(err)
    }

    for _, c := range containers {
        if (strings.Contains(c.Image, instanceKey)) {
            status := storage.Save(c)
            fmt.Println("[INFO] monitor:", status)
        }
    }
}
