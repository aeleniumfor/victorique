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

// GetContainerNameList method is docker ps
func (c *Container) GetContainerNameList() {
	containerlist := []string{}
	cli := Connect()
	containers, _ := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	for i, container := range containers {
		fmt.Println(container.Names)
		containerlist = append(containerlist, container.Names[0])
		fmt.Println(i)
	}
	fmt.Println(containerlist)
}

func main() {
	s := New()
	s.GetContainerNameList()
}
