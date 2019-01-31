package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Container type struct
type Container struct {
	id      string
	inspect types.ContainerJSON
	cli     *client.Client
}

// Containers type struct
type Containers struct {
	host      string
	container Container
}

// New Container struct
func New() (c *Container) {
	tr := &http.Transport{}
	cli, err := client.NewClient("http://localhost:2375", client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		panic(err)
	}

	return &Container{cli: cli}
}

// GetContainerNameList method is docker container name list
func (c *Container) GetContainerNameList() []string {
	containerlist := []string{}
	containers, _ := c.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	for _, container := range containers {
		containerlist = append(containerlist, container.Names[0])
	}
	return containerlist
}

// GetContainerInspect is Container inspect
func (c *Container) GetContainerInspect() {
	i, _ := c.cli.ContainerInspect(context.Background(), c.id)
	c.inspect = i
}

// CreateContainer is Container create return Container ID
func (c *Container) CreateContainer() {
	container, _ := c.cli.ContainerCreate(context.Background(), &container.Config{
		Image: "alpine",
		Tty:   true,
	}, &container.HostConfig{
		AutoRemove: true,
	}, nil, "container")
	c.id = container.ID
}

// StopContainer is docker stop <id>
func (c *Container) StopContainer() {
	time.Sleep(4)
	fmt.Println(c.id)
	c.cli.ContainerStop(context.Background(), c.id, nil)
}

// RunContainer is docker run
func (c *Container) RunContainer() {
	c.cli.ContainerStart(context.Background(), c.id, types.ContainerStartOptions{})
}

func main() {
	s := New()
	s.CreateContainer()
	s.RunContainer()
	s.GetContainerInspect()
	s.StopContainer()
}
