package storage

import (
    "bytes"
    "net/http"
    "io/ioutil"
    "encoding/json"

    "github.com/fsouza/go-dockerclient"
)

const (
    apiEndpoint = "http://service-storage.dev/api/v1/instances"
)

// Message represents data of each container in the list returned by
// ListContainers that is required by the service-storage
type Message struct {
    Id       string
    Image    string
    Command  string
    Created  int64
    State    string
    Status   string
    Ports    []docker.APIPort
    Networks docker.NetworkList
}

// Save creates PUT http request and
// pushes info about container to service
func Save(data docker.APIContainers) string {
    prepared := bytes.NewBuffer(prepareData(data))
    req, err := http.NewRequest(http.MethodPut, apiEndpoint, prepared)
    if err != nil {
        panic(err)
    }

    resp, nil := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    return string(body)
}

// PrepareData creates raw bytes array for api call
func prepareData(container docker.APIContainers) []byte {
    msg := Message{
        Id:       container.ID,
        Image:    container.Image,
        Command:  container.Command,
        Created:  container.Created,
        Status:   container.Status,
        Ports:    container.Ports,
        Networks: container.Networks,
    }

    rawBytes, err := json.Marshal(msg)
    if err != nil {
        panic(err)
    }

    return rawBytes
}
