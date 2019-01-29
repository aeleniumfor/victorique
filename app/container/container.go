package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"
)

// Container type struct
type Container struct {
	id     string
	name   string
	ip     string
	status string
}

// Containers type struct
type Containers struct {
	host      string
	container Container
}

// New Container struct
func New() (c *Container) {
	return &Container{}
}

// Connect docker server
func Connect() (cli *client.Client) {
	tr := &http.Transport{}
	cli, err := client.NewClient("http://localhost:2375", client.DefaultVersion, &http.Client{Transport: tr}, map[string]string{})
	if err != nil {
		panic(err)
	}
	return cli
}

// GetContainer method is docker ps
func (c *Container) GetContainer() {
	cli := Connect()
	containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{})
}

func main() {
	s := New()
	s.GetContainer()
}
