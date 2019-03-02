package container

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Containers is managed
type Containers struct {
	host      string
	client    *client.Client
	Container []Container
}

// Container is managed
type Container struct {
	name    string
	id      string
	inspect types.ContainerJSON
}

// New is manager
func New(hostlist []string) (c *Containers) {
	tr := &http.Transport{}
	hostname := "http://localhost:2375"
	cli, err := client.NewClient(hostname, client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		log.Println(err)
	}
	return &Containers{client: cli}
}

// GetContainerList is multi host get container list
func (c *Containers) GetContainerList() {
	containers, _ := c.client.ContainerList(context.Background(), types.ContainerListOptions{})
	fmt.Println(containers)
}
