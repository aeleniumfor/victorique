package container

import (
	"context"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Containers is managed
type Containers struct {
	host      string
	client    *client.Client
	Container []Container // slice
}

// Container is managed
type Container struct {
	Name    string
	Id      string
	Inspect types.ContainerJSON
}

// New is manager
func New() (c *Containers) {
	tr := &http.Transport{}
	hostname := "http://localhost:2376"
	cli, err := client.NewClient(hostname, client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		log.Println(err)
	}
	return &Containers{host: hostname,client: cli}
}

// GetContainerList is multi host get container list
func (c *Containers) GetContainerList() []string {
	containers, _ := c.client.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	for _, container := range containers {
		c.Container = append(c.Container, container.Names[0])
	}
	return containerList
}
