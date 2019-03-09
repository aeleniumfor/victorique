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
	Container []Container // Container List
}

// Container is managed
type Container struct {
	Name string
	ID   string
	IP   string
}

// New is manager
func New(hostname string) (c *Containers) {
	tr := &http.Transport{}
	cli, err := client.NewClient(hostname, client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		log.Println(err)
	}
	return &Containers{host: hostname, client: cli}
}

// GetContainerList is multi host get container list
func (c *Containers) GetContainerList() *Containers {
	list, _ := c.client.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	for i := 0; i < len(list); i++ {
		List := Container{
			list[i].Names[0],
			list[i].ID,
			list[i].NetworkSettings.Networks["bridge"].IPAddress, // 中身はmap
		}
		c.Container = append(c.Container, List)
	}
	return c
}

// GetContainerNameList is Get Containers Name
func (c *Containers) GetContainerNameList() []string {
	containerNameList := []string{}
	for i := 0; i < len(c.Container); i++ {
		containerNameList = append(containerNameList, c.Container[i].Name)
	}
	return containerNameList
}
